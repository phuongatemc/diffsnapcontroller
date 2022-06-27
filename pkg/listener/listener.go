package listener

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kubernetes-csi/csi-lib-utils/connection"
	"github.com/kubernetes-csi/csi-lib-utils/metrics"
	changeblockservice "github.com/phuongatemc/diffsnapcontroller/pkg/changedblockservice/changed_block_service"
	"github.com/phuongatemc/diffsnapcontroller/pkg/controller"
	"k8s.io/klog/v2"
	"net/http"
)

var (
	listenerPort string = ":8000"
)

type Listener struct {
	httpServer                     *mux.Router
	differentialSnapshotGrpcClient changeblockservice.DifferentialSnapshotClient
}

type ChangeBlocksResponse struct {
	ChangeBlockList []ChangedBlock `json:"changeBlockList"`      //array of ChangedBlock
	NextOffset      string         `json:"nextOffset,omitempty"` // StartOffset of the next “page”.
	VolumeSize      uint64         `json:"volumeSize"`           // size of volume in bytes
	Timeout         uint64         `json:"timeout"`              //second since epoch
}

type ChangedBlock struct {
	Offset  uint64 `json:"offset"`            // logical offset
	Size    uint64 `json:"size"`              // size of the block data
	Context []byte `json:"context,omitempty"` // additional vendor specific info.  Optional.
	ZeroOut bool   `json:"zeroOut"`           // If ZeroOut is true, this block in SnapshotTarget is zero out.
	// This is for optimization to avoid data mover to transfer zero blocks.
	// Not all vendors support this zeroout.
}

func NewListener(csiAddress string) (*Listener, error) {
	metricsManager := metrics.NewCSIMetricsManagerForSidecar("cbt-service")
	//create client
	csiConn, err := connection.Connect(
		csiAddress,
		metricsManager,
		connection.OnConnectionLoss(connection.ExitOnConnectionLoss()))
	if err != nil {
		return nil, err
	}

	cbtGrpcClient := changeblockservice.NewDifferentialSnapshotClient(csiConn)

	listener := &Listener{
		httpServer:                     mux.NewRouter(),
		differentialSnapshotGrpcClient: cbtGrpcClient,
	}
	return listener, nil
}

func (l Listener) StartListener() {
	l.httpServer.HandleFunc("/{cr-namespace}/{cr-name}/changedblocks", l.ServeHttpRequestHandler).Methods("GET")
	// start listening
	err := http.ListenAndServe(listenerPort, l.httpServer)
	klog.Fatalf("%v", err)
}

func (l Listener) ServeHttpRequestHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	snapshotDeltaCrName := vars["cr-name"]

	// get the start-offset embedded as a query parameter in the request URL
	startOffset := req.URL.Query().Get("startoffset")

	if _, ok := controller.DSMap[snapshotDeltaCrName]; !ok {
		klog.Errorf("Missing VolumeSnapshotDelta CR details")
		http.Error(resp, "Unable to get VolumeSnapshotDelta details", http.StatusInternalServerError)
		return
	}

	cbs, err := l.differentialSnapshotGrpcClient.GetChangedBlocks(context.TODO(), &changeblockservice.GetChangedBlocksRequest{
		SnapshotBase:   controller.DSMap[snapshotDeltaCrName].BaseVS.Name,
		SnapshotTarget: controller.DSMap[snapshotDeltaCrName].TargetVS.Name,
		StartOfOffset:  startOffset,
		MaxEntries:     controller.DSMap[snapshotDeltaCrName].MaxEntries,
	})
	if err != nil {
		klog.Errorf("Unable to get changed blocks from listener service: %v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	klog.Infof("Processed GetChangedBlocks %#v", cbs)

	var changedBlocks []ChangedBlock

	for _, cb := range cbs.ChangedBlocks {
		changedBlocks = append(changedBlocks, ChangedBlock{
			Offset:  cb.Offset,
			Size:    cb.Size,
			Context: cb.Context,
			ZeroOut: cb.ZeroOut,
		})
	}

	// Send a success response along with the payload
	respondWithJSON(resp, http.StatusOK, ChangeBlocksResponse{
		ChangeBlockList: changedBlocks,
		NextOffset:      cbs.NextOffSet,
		VolumeSize:      cbs.VolumeSize,
		Timeout:         cbs.Timeout,
	})
}

func respondWithJSON(resp http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(code)
	resp.Write(response)
}

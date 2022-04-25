FROM golang:1.18 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/diffsnap-controller .

FROM gcr.io/distroless/base
COPY --from=build /build /bin
ENTRYPOINT ["/bin/diffsnap-controller"]

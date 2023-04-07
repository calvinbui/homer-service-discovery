FROM golang:1.20.3 as builder
WORKDIR /build
COPY . .
RUN \
  CGO_ENABLED=0 \
  go build -o main cmd/main.go

FROM busybox:1.36.0
COPY --from=builder /build/main /main
ENTRYPOINT ["/main"]

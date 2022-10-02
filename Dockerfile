FROM golang:1.19.1-alpine3.16 as builder
WORKDIR /build
COPY . .
RUN \
  CGO_ENABLED=0 \
  go build -o main cmd/main.go

FROM busybox:1.34.1
COPY --from=builder /build/main /main
ENTRYPOINT ["/main"]

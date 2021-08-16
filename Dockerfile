FROM golang:1.16.7-alpine3.14 as builder
WORKDIR /build
COPY . .
RUN \
  CGO_ENABLED=0 \
  go build -o main cmd/main.go

FROM busybox:1.33.1
COPY --from=builder /build/main /main
ENTRYPOINT ["/main"]

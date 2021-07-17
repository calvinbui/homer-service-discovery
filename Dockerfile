FROM golang:1.16.5-alpine as builder
WORKDIR /build
COPY . .
RUN \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -o main cmd/main.go

FROM scratch
COPY --from=builder /build/main /main
ENTRYPOINT ["/main"]

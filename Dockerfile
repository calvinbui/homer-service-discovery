FROM golang:1.16.7-alpine as builder
WORKDIR /build
COPY . .
RUN \
  CGO_ENABLED=0 \
  go build -o main cmd/main.go

FROM scratch
COPY --from=builder /build/main /main
ENTRYPOINT ["/main"]

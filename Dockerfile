FROM golang:1.16.5-alpine as builder
WORKDIR /build
COPY . .
RUN go build -o main cmd/main.go

FROM scratch
COPY --from=builder /build/main /main
ENTRYPOINT ["./main"]

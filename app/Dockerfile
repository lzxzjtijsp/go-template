# syntax=docker/dockerfile:1

FROM golang:1.20.4 as builder
WORKDIR /go-template
COPY .. .
RUN go mod download
RUN go mod tidy
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
COPY --from=builder /go-template/main /go-template/main
CMD ["/go-template/main"]

FROM golang:latest as builder

ENV GO111MODULE=on

COPY go.mod /go/api/
COPY go.sum /go/api/

WORKDIR /go/api

RUN go mod download

COPY . /go/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main cmd/main.go

FROM alpine:latest

COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]

EXPOSE 8080
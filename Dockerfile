FROM golang:1.14.0-stretch as builder

ENV GOPROXY https://proxy.golang.org

ADD go.* ./

RUN go mod download

WORKDIR /go/src/github.com/AlfredoPastor/ari-queue

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/AlfredoPastor/ari-queue/cmd/main.go

## ----------------------------
FROM scratch

## copy ssl certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy bin
COPY --from=builder /binary /ari-queue
CMD ["/ari-queue"]

FROM golang:1.15 as builder
WORKDIR /go/src/github.com/beerproto/dataset/src/cmd/csvbeer
COPY . /go/src/github.com/beerproto/dataset/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -o csvbeer .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/beerproto/dataset/src/cmd/csvbeer .
ENTRYPOINT ["./csvbeer"]
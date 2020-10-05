FROM golang:1.15 as builder
WORKDIR /go/src/github.com/beerproto/dataset/src/cmd/csvbeer
COPY src /go/src/github.com/beerproto/dataset/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -o csvbeer .

#FROM alpine:3.12
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#COPY --from=builder /go/src/github.com/beerproto/dataset/src/cmd/csvbeer .
#ENTRYPOINT ["./csvbeer"]

COPY entrypoint.sh entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
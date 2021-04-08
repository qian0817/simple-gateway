FROM golang:1.13.8 as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build .
#
#FROM alpine
#COPY --from=builder /app/gateway ./gateway
EXPOSE 8080
CMD ["./gateway"]
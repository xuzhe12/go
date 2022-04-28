FROM golang:1.18-alpine as builder

ENV GOPATH=''

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY zuoye2.go /app/zuoye2.go

WORKDIR /app


RUN GOOS=linux GOARCH=amd64 go mod init zuoye2 && go mod tidy && go build .


FROM alpine

COPY --from=builder /app/zuoye2 /app/zuoye2


ENTRYPOINT ["/app/zuoye2"]

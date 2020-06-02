FROM golang:1.12-buster as builder
ENV GO111MODULE on
ADD . /gin-example
WORKDIR /gin-example
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server

FROM debian:buster-slim
COPY --from=builder /gin-example/server /bin/server
EXPOSE 9000
ENTRYPOINT ["/bin/server"]

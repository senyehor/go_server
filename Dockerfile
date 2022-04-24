# syntax=docker/dockerfile:1
FROM golang:1.17.7 as builder

WORKDIR /go_server
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o server ./

FROM gcr.io/distroless/base-debian11:debug
# todo think of cfg directory + env
WORKDIR /go_server
USER go_server
COPY --from=builder /go_server/server .
COPY --from=builder /go_server/prod_config.yml /bin/
COPY --from=builder /go_server/packet_config.yml /bin/
ENTRYPOINT ["/bin/server"]

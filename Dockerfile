# syntax=docker/dockerfile:1
FROM golang:1.17.7 as builder

WORKDIR /go_server
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o server ./

FROM gcr.io/distroless/base-debian11:debug
WORKDIR /packet_listener
WORKDIR ./bin
COPY --from=builder /go_server/server ./bin/

WORKDIR /packet_listener/cfg
COPY --from=builder /go_server/prod_config.yml ./cfg/
COPY --from=builder /go_server/packet_config.yml ./cfg/
COPY --from=builder /go_server/app_user_db_settings.env ./cfg/

ENV GO_APP_CONFIG_PATH="/packet_listener/cfg/"

ENTRYPOINT ["/packet_listener/bin/server"]

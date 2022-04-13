# syntax=docker/dockerfile:1
FROM golang:1.17.7 as builder

WORKDIR /go_server
COPY src/github.com/senyehor/go_server .
COPY prod_config.yml .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o server ./

FROM gcr.io/distroless/base-debian11:debug

COPY --from=builder /go_server/server /bin/
COPY --from=builder /go_server/prod_config.yml /bin/
ENTRYPOINT ["/bin/server"]


FROM --platform=linux/amd64 kong/go-plugin-tool:2.0.4-alpine-latest AS builder

RUN mkdir -p /tmp/plugin/

COPY . /tmp/plugin/

RUN cd /tmp/plugin/ && \
    go get github.com/Kong/go-pdk && \
    go mod init kong-go-plugin && \
    go build -buildmode plugin plugin/go-hello.go

FROM --platform=linux/amd64 kong:2.8.1-alpine

COPY --from=builder /tmp/plugin/go-hello.so /usr/local/bin

USER kong

#!/bin/bash

docker build -t kong-with-plugin-image:local -f plugin/Dockerfile .
docker run -ti --rm --name kong-with-plugin \
  -e "KONG_DATABASE=off" \
  -e "KONG_PLUGINSERVER_NAMES=/usr/local/bin" \
  -e "KONG_PLUGINSERVER_GO_HELLO_SOCKET=/usr/local/kong/go-hello.socket" \
  -e "KONG_PLUGINSERVER_GO_HELLO_START_CMD=/usr/local/bin/go-hello" \
  -e "KONG_PLUGINSERVER_GO_HELLO_QUERY_CMD=/usr/local/bin/go-hello -dump" \
  -e "KONG_PLUGINS=bundled" \
  -e "KONG_PROXY_LISTEN=0.0.0.0:8000" \
  -p 8000:8000 \
  kong-with-plugin-image:local

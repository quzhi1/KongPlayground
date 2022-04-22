# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')

# Kong
k8s_yaml('all-in-one-postgres.yaml')

kong_resource_map = {
  "ingress-kong": [31104],
  "postgres": [5432],
  "kong-migrations": [],
}

for kong_resource, ports in kong_resource_map.items():
  k8s_resource(
    kong_resource,
    port_forwards=ports,
    labels="kong",
  )

local_resource(
  'expose-kong-proxy',
  '',
  serve_cmd='kubectl -n kong port-forward service/kong-proxy 8080:80',
  resource_deps=["ingress-kong"],
  labels="kong",
)

local_resource(
  'expose-kong-ingress',
  '',
  serve_cmd='kubectl -n kong port-forward service/kong-validation-webhook 8444:8444',
  resource_deps=["ingress-kong"],
  labels="kong",
)

# Sync Kong config
local_resource(
  "apply-kong-config",
  "deck sync --kong-addr https://localhost:8444 --tls-skip-verify",
  deps=["kong.yaml"],
  resource_deps=["expose-kong-ingress"],
  labels="kong",
)

# Example application
compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/hello-world example-application/main.go'

local_resource(
  'hello-world-compile',
  compile_cmd,
  deps=['example-application/main.go'],
  labels="example-application",
)

docker_build_with_restart(
  'hello-world-image',
  '.',
  entrypoint=['/opt/app/bin/hello-world'],
  dockerfile='example-application/Dockerfile',
  only=[
    './bin',
  ],
  live_update=[
    sync('./bin', '/opt/app/bin'),
  ],
)

k8s_yaml(helm('example-application/helm'))

k8s_resource(
  "hello-world",
  port_forwards=8090,
  labels="example-application",
)

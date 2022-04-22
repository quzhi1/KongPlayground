# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')

# Kong
helm_remote(
  'kong',
  repo_name='kong',
  repo_url='https://charts.konghq.com',
  values="minimal-kong-standalone.yaml"
)

kong_resource_map = {
  "kong-kong": [30933, 8000, 8001, 8002, 8444],
  "kong-postgresql": [5432],
  "kong-kong-init-migrations": [],
  "kong-kong-post-upgrade-migrations": [],
  "kong-kong-pre-upgrade-migrations": [],
}

for kong_resource, ports in kong_resource_map.items():
  k8s_resource(
    kong_resource,
    port_forwards=ports,
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

# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')
compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

# Compile Kong go plugin
local_resource(
  'kong-plugin-compile',
  # 'go build -buildmode plugin -o bin/go-hello.so plugin/go-hello.go',
  compile_opt + 'go build -o bin/go-hello plugin/go-hello.go',
  # 'go build -o bin/go-hello plugin/go-hello.go',
  deps=['plugin/go-hello.go'],
  labels="kong",
)

# Build Kong ingress with plugin binary
docker_build(
  'kong-with-plugin-image',
  '.',
  dockerfile='plugin/Dockerfile',
  only=[
    'bin/go-hello',
  ],
)

# Deploy Kong
# k8s_yaml('all-in-one-postgres.yaml')
helm_remote(
  'kong',
  repo_name='kong',
  repo_url='https://charts.konghq.com',
  values="developer.kong.values.yaml",
)

# Kong services config
kong_resource_map = {
  # "ingress-kong": [],
  "kong-kong": [],
  "kong-postgresql": 5432,
  # "kong-migrations": [],
}

# Lable Kong and port forward
for kong_resource, ports in kong_resource_map.items():
  k8s_resource(
    kong_resource,
    port_forwards=ports,
    labels="kong",
  )

# Port forward kong-proxy
local_resource(
  'expose-kong-proxy',
  '',
  serve_cmd='kubectl port-forward service/kong-kong-proxy 8080:80',
  resource_deps=["kong-kong"],
  labels="kong",
)

# Port forward kong-ingress (admin)
local_resource(
  'expose-kong-ingress',
  '',
  serve_cmd='kubectl port-forward service/kong-kong-admin 8444:8444',
  resource_deps=["kong-kong"],
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

# Compile example application
local_resource(
  'hello-world-compile',
  compile_opt + 'go build -o bin/hello-world example-application/main.go',
  deps=['example-application/main.go'],
  labels="example-application",
)

# Build example docker image
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

# Install example helm chart
k8s_yaml(helm('example-application/helm'))

# Label and port forwarding example applciation
k8s_resource(
  "hello-world",
  port_forwards=8090,
  labels="example-application",
)

# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')
compile_opt = 'GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 '

# Compile Kong go plugin
local_resource(
  'kong-plugin-compile',
  compile_opt + 'go build -o bin/go-hello plugin/go-hello.go',
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
helm_remote(
  'kong',
  repo_name='kong',
  repo_url='https://charts.konghq.com',
  values="developer.kong.values.yaml",
)

# Lable postgresql and port forward
k8s_resource(
  "kong-postgresql",
  port_forwards=5432,
  labels="kong",
)

# Init migration
k8s_resource(
  "kong-kong-init-migrations",
  labels="kong",
  resource_deps=["kong-postgresql"],
)

# Pre upgrade migration
k8s_resource(
  "kong-kong-pre-upgrade-migrations",
  labels="kong",
  resource_deps=["kong-kong-init-migrations"],
)

# Post upgrade migration
k8s_resource(
  "kong-kong-post-upgrade-migrations",
  labels="kong",
  resource_deps=["kong-kong-pre-upgrade-migrations"],
)

# Lable Kong and port forward
k8s_resource(
  "kong-kong",
  port_forwards=[8000, 8001],
  labels="kong",
  resource_deps=["kong-kong-post-upgrade-migrations"],
)

# Sync Kong config
local_resource(
  "apply-kong-config",
  "deck sync --kong-addr http://localhost:8001",
  deps=["kong.yaml"],
  resource_deps=["kong-kong"],
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

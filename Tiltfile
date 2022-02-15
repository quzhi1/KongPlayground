# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')

helm_remote(
  'kong',
  repo_name='kong',
  namespace='kong',
  repo_url='https://charts.konghq.com',
  values="full-k4k8s-with-kong-enterprise.yaml"
)

k8s_resource(
  "kong-kong",
  port_forwards=[30933, 8000, 8001, 8002, 8444],
)

k8s_resource(
  "kong-postgresql",
  port_forwards=5432,
)
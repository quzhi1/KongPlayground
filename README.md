# KongPlayground
Playground for Kong API gateway

## Install with Helm
```bash
tilt up
```

## Testing
1. Use Postman to call http://localhost:8080/hello
2. It should redirect to Postman Echo site

## Trouble shooting
- Kong config: https://localhost:8444
- Kong services: https://localhost:8444/services
- Kong routes: https://localhost:8444/routes
- Kong plugins: https://localhost:8444/plugins

## Kong k8s yaml
I got the yaml file from Kong's Github repo: https://github.com/Kong/kubernetes-ingress-controller/blob/main/deploy/single/all-in-one-postgres.yaml

There are several changes you need to make:

### Expose admin API
In order to talk to admin API, I changed the configuration of service `kong-validation-webhook`:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: kong-validation-webhook
  namespace: kong
spec:
  ports:
  - name: webhook
    port: 443
    protocol: TCP
    targetPort: 8080
  - name: admin-url
    port: 8444
    protocol: TCP
    targetPort: 8444
  type: LoadBalancer
  selector:
    app: ingress-kong
```

### Replace official image
You also need to replace official Kong image with an image with go plugin. So change the image name from `image: kong:2.8` to `image: kong-with-plugin-image`

## Set environment variable
Find the `proxy` container, and add the following environment variables:
```yaml
- name: KONG_PLUGINSERVER_NAMES
  value: go-hello
- name: KONG_PLUGINSERVER_GO_HELLO_SOCKET
  value: /usr/local/kong/go-hello.socket
- name: KONG_PLUGINSERVER_GO_HELLO_START_CMD
  value: /usr/local/bin/go-hello
- name: KONG_PLUGINSERVER_GO_HELLO_QUERY_CMD
  value: /usr/local/bin/go-hello -dump
- name: KONG_PLUGINS
  value: bundled,go-hello
```
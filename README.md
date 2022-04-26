# KongPlayground
Playground for Kong API gateway

## Install with Helm
```bash
tilt up
```

## Testing
1. Use Postman to call http://localhost:8080/hello
2. Find following headers:
   1. `kong-hello-world-unique-id` is set by correlation-id plugin. This plugin is provided by Kong.
   2. `x-hello-from-go` is set by a custom go plugin.

## Trouble shooting
- Kong config: https://localhost:8444
- Kong services: https://localhost:8444/services
- Kong routes: https://localhost:8444/routes
- Kong plugins: https://localhost:8444/plugins

## Note
### Kong k8s yaml
I got the yaml file from Kong's Github repo: https://github.com/Kong/kubernetes-ingress-controller/blob/main/deploy/single/all-in-one-postgres.yaml
However, in order to talk to admin API, I changed the configuration of service `kong-validation-webhook`:
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
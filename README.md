# KongPlayground
Playground for Kong API gateway

## Install with Helm
```bash
tilt up
```

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
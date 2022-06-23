# KongPlayground
Playground for Kong API gateway

The plugin does three things:
1. Change the routing target to mockbin.org:80.
2. Add request headers.
3. Set query parameters.

## Install with Helm
```bash
tilt up
```

## Testing
1. Use Postman to call http://localhost:8080/hello.
2. It should route to mockbin.org:80, and have new request headers and query parameters.

## Trouble shooting
- Kong config: https://localhost:8444
- Kong services: https://localhost:8444/services
- Kong routes: https://localhost:8444/routes
- Kong plugins: https://localhost:8444/plugins

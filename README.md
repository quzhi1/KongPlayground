# KongPlayground
Playground for Kong API gateway

## Install with Helm
1. Put a license file in repository root.
2. Run
```bash
kubectl create namespace kong
kubectl create secret generic kong-enterprise-license --from-file=./license -n kong
kubectl create secret generic kong-enterprise-superuser-password \
-n kong \
--from-literal=password=root
echo '{"cookie_name":"admin_session","cookie_samesite":"off","secret":"<your-password>","cookie_secure":false,"storage":"kong"}' > admin_gui_session_conf
echo '{"cookie_name":"portal_session","cookie_samesite":"off","secret":"<your-password>","cookie_secure":false,"storage":"kong"}' > portal_session_conf
kubectl create secret generic kong-session-config \
-n kong \
--from-file=admin_gui_session_conf \
--from-file=portal_session_conf
tilt up
```
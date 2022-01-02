# Deploying multiapp kc cluster with ingress

enable ingress `minikube addons enable ingress`

- pulls the `ingress-nginx` image from GCR.
- starts the ingress pods in `ingress-nginx` namespace

Created a deployment for the first and second app. pod spec has the container image, port, labels which will be used to select the pods from the service

Created two services. service selects the pods by the labels. no type specified, so default goes to `ClusterIP` service. with this type of abstraction, the nodes are not exposed outside. so `NodePort` service is not needed.

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: apps-ingress
  annotations:
    kubernetes.io/ingress.class: "nginx"
  labels:
    name: apps-ingress
spec:
  rules:
    - host: myhost.com #usually the real domain name of the node in k8 cluster. how this works in minikube?
      http:
        paths:
          - pathType: Prefix
            path: "/anything"
            backend:
              service:
                name: first-service
                port:
                  number: 80
          - pathType: Prefix
            path: "/uuid"
            backend:
              service:
                name: second-service
                port:
                  number: 80
```

edited `/etc/hosts` inside the minikube node to take `myhost.com` as the hostname for the minikube IP
then creatd the ingress object.
able to send `curl myhost.com/anything` and `curl myhost.com/uuid`. the former endpoint was sent
to the first service. the later was sent to second service.

then added an environment variable in the pods to enable application logging.
viewed the logs using `kc logs -l app=first-app`

things to follow

- what is httpbin app?
- what is GUNICORN? is it a logging app?
- more on `ingress` controller in kubernetes
- what is `ingress-nginx` image? what does it do?

```console
rhyme@ip-172-31-230-229:~/ingress-k8s$ kc get pods -n ingress-nginx
NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create--1-4j5ph     0/1     Completed   0          68m
ingress-nginx-admission-patch--1-rxxbm      0/1     Completed   1          68m
ingress-nginx-controller-69bdbc4d57-sstj8   1/1     Running     0          68m
```

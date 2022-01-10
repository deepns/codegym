# Learning to deploy apps in K8s cluster with minikube

## Steps followed

- Installed Virtual box
- Installed minikube through HomeBrew(instructions from [here](https://minikube.sigs.k8s.io/docs/start/))
- Installed docker locally, pointed docker to the daemon inside minikube

```text
~ % docker images
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
~ % eval $(minikube -p minikube docker-env)
~ % docker images
REPOSITORY                                TAG       IMAGE ID       CREATED          SIZE
testapp                                   v1        4d780fcbe334   11 minutes ago   141MB
<none>                                    <none>    069dd7258441   58 minutes ago   141MB
nginx                                     latest    605c77e624dd   2 days ago       141MB
k8s.gcr.io/kube-apiserver                 v1.22.3   53224b502ea4   2 months ago     128MB
k8s.gcr.io/kube-controller-manager        v1.22.3   05c905cef780   2 months ago     122MB
k8s.gcr.io/kube-scheduler                 v1.22.3   0aa9c7e31d30   2 months ago     52.7MB
k8s.gcr.io/kube-proxy                     v1.22.3   6120bd723dce   2 months ago     104MB
kubernetesui/dashboard                    v2.3.1    e1482a24335a   6 months ago     220MB
k8s.gcr.io/etcd                           3.5.0-0   004811815584   6 months ago     295MB
kubernetesui/metrics-scraper              v1.0.7    7801cfc6d5c0   6 months ago     34.4MB
k8s.gcr.io/coredns/coredns                v1.8.4    8d147537fb7d   7 months ago     47.6MB
gcr.io/k8s-minikube/storage-provisioner   v5        6e38f40d628d   9 months ago     31.5MB
k8s.gcr.io/pause                          3.5       ed210e3e4a5b   9 months ago     683kB
~ %
```

- Created the docker image for the app with [Dockerfile](Dockerfile) (`docker build -t testapp:v1 .`)
- Created the deployment and service spec files.
  - Specified the container image in the deployment spec.
  - Exposed the container port 80.
  - Created a NodePort service

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: testapp
spec:
  selector:
    matchLabels:
      app: testapp
  template:
    metadata:
      labels:
        app: testapp
    spec:
      containers:
        - name: testapp
          image: testapp:v1
          ports:
            - containerPort: 80
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: testapp
spec:
  selector:
    app: testapp
  type: NodePort
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

- Create the deployment and service (`kubectl create -n default -f k8s` - all yaml files under k8s/ directory are processed).

```console
testapp % kc create -n default -f k8s
deployment.apps/testapp created
service/testapp created
testapp % kc get pods
NAME                      READY   STATUS    RESTARTS   AGE
testapp-dbc7c6bbf-4sxsj   1/1     Running   0          4s
testapp %
```

- get minikube ip. get URL for services running inside minikube. `minikube service testapp` launches the service's URL in the default browser.

```console
testapp % minikube ip
192.168.59.100
testapp % minikube service testapp --url
http://192.168.59.100:30232
```

- `minikube dashboard` launches the dashboard in the browser.
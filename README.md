# develop, build & deploy go app to [k3s](https://k3s.io/)


#### Prerequisites

* Installed Docker to build and push an image to registry.
* [Installed and configured kubernetes](https://github.com/bobanbojkovski/k3s) - [k3s](https://k3s.io/) cluster. 


#### GOlang sample app to print out pod's hostname, ip and a text using [echo http server](https://echo.labstack.com/)

go_echo.go snippet:
```
...
var (
		text = flag.String("text", "default", "text to print out")
		port = flag.String("port", ":1323", "http port")
	)
...
e := echo.New()
	e.GET("/"+*text, func(c echo.Context) error {
		return c.String(http.StatusOK, *text+"\n"+hostname+"\n"+hostip)
	})
...
```

**text** and **port** can be defined as container arguments in pod deployment.

bar_deploy_svc.yaml snippet
```
spec:
      containers:
      - name: bar-app
        image: localhost:5000/echo:1.0
        args:
        - "-text=bar"
```

#### Use Docker to build image and store it in Docker local registry

Image build instructions are documented in Dockerfile.
```
docker build -t echo:1.0 .
docker tag echo:1.0 localhost:5000/echo:1.0
docker push localhost:5000/echo:1.0
```

#### Create app namespace, ingress, pod deplyoments & services

```
kubectl apply -f app_ns.yaml
kubectl apply -f ingress_app.yaml -n app
kubectl apply -f bar_deploy_svc.yaml -n app
kubectl apply -f default_deploy_svc.yaml -n app
```

#### Site access

```
http://<fqdn>/bar
http://<fqdn>/default

curl -s "http://<fqdn>/default?[1-3]" -w "\n\n"
default
default-app-deployment-84446dd48b-sqsxd
10.42.0.149

default
default-app-deployment-84446dd48b-xcmvz
10.42.0.147

default
default-app-deployment-84446dd48b-67vqj
10.42.0.148
```


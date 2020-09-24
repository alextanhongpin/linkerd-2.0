# linkerd-2.0

This repository is meant to demonstrate the capability of linkerd as well as exploring the capabilities when used with Kubernetes.

Environment:
- Docker Desktop Version 2.3.0.5

## Installation

```
$ brew install linkerd
```

Checking the version:
```
$ linkerd version
```

Output:
```bash
Client version: stable-2.8.1
Server version: stable-2.8.1
```

Validate your Kubernetes cluster:
```
$ linkerd check --pre
```

Install Linkerd onto the cluster:

```bash
$ linkerd install | kubectl apply -f -

# Fresh install.
$ linkerd install --ignore-cluster | kubectl delete -f -
```

Validate the installation:
```bash
$ linkerd check
```

View dashboard:

```
$ linkerd dashboard &
```


## Uninstall linkerd

```bash
$ linkerd uninstall | kubectl delete -f -
```

## Deploying Services

Let's deploy a server.

```bash
$ cat go-server/deployment.yaml | linkerd inject - | kubectl apply -f -
```

Verify that the server is deployed:

```bash
$ linkerd stats deployment
```

Output:

```
NAME                   MESHED   SUCCESS   RPS   LATENCY_P50   LATENCY_P95   LATENCY_P99   TCP_CONN
go-server-deployment      1/1         -     -             -             -             -          -
```

Since we are running on Docker Desktop, we will get the NodePort and call the endpoint:


```bash
$ alias k=kubectl
$ k get svc
```

Output:
```
NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
go-server-service   NodePort    10.100.238.146   <none>        80:30918/TCP   5m53s
kubernetes          ClusterIP   10.96.0.1        <none>        443/TCP        12m
```

Call the endpoint:
```bash
$ curl localhost:30918
```

## Running client-server setup

We want to deploy a pair of client/server setup, with the client making a http request to the server. We then run the command to inject linkerd mesh to both containers and see how they interact with one another.
```
# Build the docker images.
$ make build-client
$ make build-server

# Inject linkerd.
$ make inject-server
$ make-inject-client
```

# linkerd-2.0

This repository is meant to demonstrate the capability of linkerd as well as exploring the capabilities when used with Kubernetes.

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

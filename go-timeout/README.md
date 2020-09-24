# Configuring Retry

We will deploy a server, and an application with timeout configured that will call that server. We can trigger timeout by calling `curl localhost:<server-port>/?timeout=true`. It must have a `/` slash, since we configured it in the `ServiceProfile`.

	NOTE: The retry and timeout has to be applied on the server that the client is calling. In this case, it is `go-server`.

```bash
# Build and deploy the server
$ make build-server
$ make inject-server

# Build and deploy the timeout-app
$ make build-timeout
$ make inject-timeout
```

Find the `timeout-service` `NodePort` IP:

```bash
$ alias k=kubernetes
$ k get svc
```

Output:

```
NAME                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
go-retry-service     NodePort    10.100.191.112   <none>        80:31332/TCP   12m
go-server-service    NodePort    10.109.146.58    <none>        80:30503/TCP   12m
go-timeout-service   NodePort    10.103.201.83    <none>        80:32305/TCP   12m
kubernetes           ClusterIP   10.96.0.1        <none>        443/TCP        77m
```

Alternatively:
```bash
# To preview the structure as json.
$ k get svc/go-timeout-service -o json

# To extract specific fields with jsonpath.
$ k get svc/go-timeout-service -o=jsonpath='{.spec.ports[0].nodePort}'
```

Once we have the `NodePort`, we can make a call to the server and see the timeout in action.
```bash
$ curl localhost:$(k get svc/go-timeout-service -o=jsonpath='{.spec.ports[0].nodePort}')/\?timeout=true
```

To see the timeout rates:
```
$ linkerd routes deploy/go-timeout-deployment --to svc/go-server-service -o wide
```

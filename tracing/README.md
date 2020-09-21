# Distributed Tracing

We can easily add distributed tracing. Run the following command in this folder:

```bash
$ linkerd upgrade --addon-config config.yaml | k apply -f -
```

If we check the linkerd namespace, we should see that the jaeger container should be running:

```bash
$ k get svc -n linkerd
```

Output:
```
NAME                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)               AGE
linkerd-collector        ClusterIP   10.109.17.140    <none>        55678/TCP,9411/TCP    4m12s
linkerd-controller-api   ClusterIP   10.107.207.198   <none>        8085/TCP              46m
linkerd-dst              ClusterIP   10.99.30.87      <none>        8086/TCP              46m
linkerd-grafana          ClusterIP   10.100.161.208   <none>        3000/TCP              46m
linkerd-identity         ClusterIP   10.100.7.179     <none>        8080/TCP              46m
linkerd-jaeger           ClusterIP   10.101.68.174    <none>        14268/TCP,16686/TCP   4m11s
linkerd-prometheus       ClusterIP   10.102.188.233   <none>        9090/TCP              46m
linkerd-proxy-injector   ClusterIP   10.102.120.20    <none>        443/TCP               46m
linkerd-sp-validator     ClusterIP   10.97.65.183     <none>        443/TCP               46m
linkerd-tap              ClusterIP   10.100.111.220   <none>        8088/TCP,443/TCP      46m
linkerd-web              ClusterIP   10.97.108.206    <none>        8084/TCP,9994/TCP     46m
```

Otherwise, we can check if the Jaeger container is already running with the following commands:

```bash
$ k -n linkerd rollout status deploy/linkerd-collector
$ k -n linkerd rollout status deploy/linkerd-jaeger
```

The output can be seen below:
```bash
$ k -n linkerd rollout status deploy/linkerd-jaeger
```

# References
- https://linkerd.io/2/tasks/distributed-tracing/

# kubectl-design

kubectl-design helps to generate kubernetes resource yaml definations using
command line options and opens it in your editor to edit it and use it.

A simple tool to save your time.

__version compatiblity : kubectl v1.18.x__

#### Installation:

```
go get -u github.com/kunalkushwaha/kubectl-design
```



Just replace _create/run_ with`design` in kubectl command
#### Usage:
```
kubectl design [resource-name] [options]

$ kubectl design deploy test --image busybox
```
![](demo.gif)


Your feedback is welcome, feel free to create issues and PRs.

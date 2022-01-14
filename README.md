# k8s自定义对象资源实践（CRD）

学习《深入剖析Kubernetes》自定义API资源相关代码

## 开始
> go version >= 1.13

如果你开启了GO111MODULE=on,并且想要自己生成代码，你可以执行下面的操作。
```shell script
$ go mod tidy
$ go mod vendor
# 自动生成代码
$ ./hack/update-codegen.sh
```

生成的代码是在当前目录下按照**github.com/happlins/....**的目录结构，你需要将相对应的apis目录下的文件和generated目录下的文件，拷贝到pkg目录下，并删除github.com目录


克隆项目到本地
```shell script
$ git clone https://github.com/happlins/k8s-controller-custom-resource.git
$ cd k8s-controller-custom-resource
```

编译项目
```shell script
$ go build -o samplecrd-controller .
$ ./samplecrd-controller -kubeconfig=$HOME/.kube/config -alsologtostderr=true
```
您还可以使用samplecrd-controller创建一个Deployment并在Kubernetes中运行它。
请注意，在这种情况下，您无需在CMD中指定-kubeconfig(如果你是通过pod来运行的话，不指定，则默认使用挂载卷中的认证信息)，因为将使用默认的InClusterConfig。

## 使用
开始创建 Network CRD资源
```shell script
$ kubectl apply -f crd/network.yaml
```

然后，通过创建Network API实例来触发事件：
```shell script
$ kubectl apply -f example/example-network.yaml
```

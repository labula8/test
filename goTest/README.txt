#goTest 环境搭建

```
cd install

wget https://storage.googleapis.com/golang/go1.12.linux-amd64.tar.gz

tar -xzvf ./go1.9.linux-amd64.tar.gz

vim /etc/profile

export GOTEST=/data/git/labula8/test/goTest
export GOROOT=$GOTEST/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export GOPATH=$GOTEST/work


使其生效
#source /etc/profile

查看是否配置成功
# go version
go version go1.9.1 linux/amd64

cd work
mddir hello
vim hello.go

package main
import "fmt"
func main() {
    fmt.Println("Hello, 世界")
}
# go run hello.go
Hello, 世界

go run hello.go 
```
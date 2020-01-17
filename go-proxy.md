- https://goproxy.io
- https://athens.azurefd.net
- https://goproxy.cn
- https://gocenter.io
- https://mirrors.aliyun.com/goproxy/

测试goproxy
```
time GOPATH=/tmp/throw GO111MODULE=on GOPROXY=https://goproxy.io go get github.com/go-sql-driver/mysql
time GOPATH=/tmp/throw GO111MODULE=on GOPROXY=https://athens.azurefd.net go get github.com/go-sql-driver/mysql
time GOPATH=/tmp/throw GO111MODULE=on GOPROXY=https://goproxy.cn go get github.com/go-sql-driver/mysql
time GOPATH=/tmp/throw GO111MODULE=on GOPROXY=https://gocenter.io go get github.com/go-sql-driver/mysql
time GOPATH=/tmp/throw GO111MODULE=on GOPROXY=https://mirrors.aliyun.com/goproxy/ go get github.com/go-sql-driver/mysql
```
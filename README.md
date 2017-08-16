# QiniuDataMigration
## 七牛云跨账户数据自动迁移


先将七牛相关go依赖安装好：
```bash
$ go get -u github.com/qiniu/api.v7
```
可以参考以下路径：
https://developer.qiniu.com/kodo/sdk/1238/go#rs-fetch


在安装api.v7时会有错误产生，如下：
C:\GOPATH>go get golang.org/x/net/context
package golang.org/x/net/context: unrecognized import path "golang.org/x/net/con
text" (https fetch: Get https://golang.org/x/net/context?go-get=1: dial tcp 216.
239.37.1:443: connectex: A connection attempt failed because the connected party
 did not properly respond after a period of time, or established connection fail
ed because connected host has failed to respond.)

原因是golang.org访问时提示Nothing to see here; move along.即官方链接更新，只需要下载安装：
```bash
$ go get github.com/golang/net
```
然后将github.com的名称改为golang.org/x/net/context 即可。

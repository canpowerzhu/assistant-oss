[TOC]
# 背景
此工具主要用于线上服务器上传下载大文件。借助OSS作为中转;默认大于500MB的文件进行分片上传。

# 基于windows构建
## windows
CGO_ENABLED=0 
GOOS=windows 
GOARCH=amd64 
go build -o assistant-oss.exe main.go
## linux
set CGO_ENABLED=0 
set GOOS=linux 
set GOARCH=amd64 
go build -o assistant-oss main.go


# 执行
```shell
[root@localhost application]# ./assistant-oss --uploadpath /application/nexus.tar.gz 
需要操作的是一个文件nexus.tar.gz; 路径是：/application/
文件/application/nexus.tar.gz分片上传
分片数量 63
Transfer Started, ConsumedBytes: 0, TotalBytes 536972042.
Transfer Data, ConsumedBytes: 536972042, TotalBytes 536972042, 100%.
Transfer Completed, ConsumedBytes: 536972042, TotalBytes 536972042.
Transfer Started, ConsumedBytes: 0, TotalBytes 536972042.
Transfer Data, ConsumedBytes: 34045952, TotalBytes 536972042, 6%.^C
```
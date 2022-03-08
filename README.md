## ipaddress


写的第一份go代码挺渣的
修改至项目地址：https://gitee.com/lionsoul/ip2region


#### 使用方法

单个IP返回查询结果

./ipadds 1.1.1.1

读入一个文件，将结果输出到控制台

./ipadds -f iplist.txt

读入一个文件，并将结果写入到一个文件

./ipadds -f iplist.txt -o /tmp/ip

##### 参数

-h 同--help

-f [ipfile] 读本地文件一行一个IP

-o [write file] 写入本地文件地址，默认输出控制台

##### webapi调用方法


启动服务：./ipaddweb db 80 #文件 db文件 端口

浏览器：http://127.0.0.1/?ip=1.1.1.1

```bash
-o 需要输出的文件名
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ipaddweb_mac web.go #编译mac架构
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ipaddweb_win web.go #编译win
go build -o ipaddweb web.go
```

======================================================

##### 文件结构

`ipadds`      依赖库

`db`          本地数据库文件，可以在原作者库内下载到最新库

`ipadds.zip`  编译后的各平台的可执行文件，使用方法详情 [使用方法](#使用方法) 

`main.go`     程序的入口文件


```bash
-o 需要输出的文件名
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ipadds_mac main.go #编译mac架构
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ipadds_win main.go #编译win
go build -o ipadds main.go
```

#### 写在最后

测试90w条IP：输出控制台查询9.6s，写本地文件4.2s

```bash
┌──(kali㉿localhost)-[~/tools]
└─$ ./ipadds -f ../源码/golang/ipadd/iplist.txt 
...
数据：930050行 耗时：9.588844719s                                                                     

┌──(kali㉿localhost)-[~/tools]
└─$ ./ipadds -f ../源码/golang/ipadd/iplist.txt -o /tmp/12333
数据：930050行 耗时：4.177968057s 
```

老板来个小星星呗 :bulb:

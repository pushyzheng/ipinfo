# ipinfo
Query IP information API for Go

## Install

Firstly,  ensure your Go version is above 1.12, then build `ipinfo.go`, you can install in Linux:

```shell
$ sudo go build ipinfo.go
$ sudo mv ipinfo /usr/bin
```

In Windows, please add `ipinfo-master` add your environment variable.

## Usage

Query the location of IP：

```shell
# query local machine IP location
$ ipinfo myip
or 
$ ipinfo

# query specified IP location
$ ipinfo 220.181.38.148
```

Query the domain name bound by IP：

```shell
$ ipinfo -r 220.181.38.148
220.181.38.148 上的网站，绑定过的域名如下：

ipv6.baidu.com                       2019-05-26-----2019-12-31
baidu.com                            2019-05-30-----2019-12-31
bdimg.com                            2019-06-02-----2019-12-31
......
```


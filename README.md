# ipinfo
Querying IP information API for terminal.

## Install

Firstly,  make ensure your Go version is above 1.12, then build `ipinfo.go` file. So you can install in Linux:

```shell
$ sudo go build ipinfo.go
$ sudo mv ipinfo /usr/bin
```

In Windows, you need to add the path of `ipinfo` to your environment variables after compilation.

## Usage

Querying the location of the IP:

```shell
# querying the local machine IP location
$ ipinfo myip
or 
$ ipinfo

# querying the specified IP location
$ ipinfo 220.181.38.148
```

Querying the domain name bound to the IP:

```shell
$ ipinfo -r 220.181.38.148
220.181.38.148 上的网站，绑定过的域名如下：

ipv6.baidu.com                       2019-05-26-----2019-12-31
baidu.com                            2019-05-30-----2019-12-31
bdimg.com                            2019-06-02-----2019-12-31
......
```


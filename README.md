# FTP Server

Simple FTP Server written in Go.

## Run Server

```sh
$ go run ./main.go
```

## Connect with FTP Client

Use [`ftp`](https://ftp.netbsd.org/pub/NetBSD/misc/tnftp/) command:

```sh
$ ftp username@localhost 8000
```


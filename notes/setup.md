# Version and Upgrade

```shell
(base) yanjun@chinablue:~$ which go
/usr/local/go/bin/go
(base) yanjun@chinablue:~$ go version
go version go1.15.1 linux/amd64
```

```shell
(base) yanjun@chinablue:~/Downloads$ wget https://go.dev/dl/go1.18.1.linux-386.tar.gz
--2022-04-17 13:13:25--  https://go.dev/dl/go1.18.1.linux-386.tar.gz

# Verify SHA256 checksum
(base) yanjun@chinablue:~/Downloads$ sha256sum ./go1.18.1.linux-386.tar.gz 
9a8df5dde9058f08ac01ecfaae42534610db398e487138788c01da26a0d41ff9  ./go1.18.1.linux-386.tar.gz

# Extract the files
(base) yanjun@chinablue:~/Downloads$ sudo tar -xvf ./go1.18.1.linux-386.tar.gz
(base) yanjun@chinablue:~/Downloads$ sudo mv go /usr/local/

(base) yanjun@chinablue:~/Downloads$ which go
/usr/local/go/bin/go
(base) yanjun@chinablue:~/Downloads$ go version
go version go1.18.1 linux/386
```
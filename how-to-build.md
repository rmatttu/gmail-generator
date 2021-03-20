# How to build

Build in Ubuntu 20.04

## Install

* ビルドには go v1.4 が必要
* GOROOT に書き込める権限が必要（？）

Go 1.16 をインストール

```bash
cd ~
wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz
mkdir tmp
tar xf go1.16.2.linux-amd64.tar.gz -C tmp
mv tmp go1.6
rm -r tmp
```

Go 1.4 をインストール

```bash
cd ~
wget https://golang.org/dl/go1.4.3.linux-amd64.tar.gz
mkdir tmp
tar xf go1.4.3.linux-amd64.tar.gz -C tmp
mv tmp/go go1.4
rm -r tmp
```

環境変数の追加

```bash
echo 'export GOROOT=$HOME/go1.16/go/' >> ~/.bash_profile
echo 'export GOTOOLDIR=$HOME/go1.16/go/pkg/tool/linux_amd64' >> ~/.bash_profile
echo 'export GOPATH=$HOME/go' >> ~/.bash_profile
echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.bash_profile
```

## Build

```bash
GOOS=windows GOARCH=amd64 go build sample.go
GOOS=darwin GOARCH=amd64 go build sample.go
GOOS=linux GOARCH=amd64 go build sample.go
```

```txt
$GOOS     $GOARCH
darwin    386
darwin    amd64
freebsd   386
freebsd   amd64
freebsd   arm
linux     386
linux     amd64
linux     arm
netbsd    386
netbsd    amd64
netbsd    arm
openbsd   386
openbsd   amd64
plan9     386
plan9     amd64
windows   386
windows   amd64
nacl      amd64
nacl      386
```

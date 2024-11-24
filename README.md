- Proxy type HTTP

- Cài đặt golang

```
- Ubuntu : 
apt update -y
apt upgrade -y
apt install -y screen
apt install -y zmap
sudo apt-get install curl git mercurial make binutils bison gcc build-essential -y

curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
source ~/.gvm/scripts/gvm

gvm install go1.20
gvm use go1.20 --default
```

```
- Centos
yum update -y
yum upgrade -y
yum install -y screen
yum install -y zmap
curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
source ~/.gvm/scripts/gvm
gvm install go1.20
gvm use go1.20 --default
```

```
- Nếu lỗi cài Golang thì có thể cài 1 bản cũ sau đó cài bản mới
gvm install go1.17.13
gvm use go1.17.13 --default

gvm install go1.20
gvm use go1.20 --default
```

- Usage : 
- go build scan.go
- zmap -p8080 -q | ./scan 8080  (thay 8080 thành port cần scan)

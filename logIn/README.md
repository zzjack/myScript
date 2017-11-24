# LogIn
It is convenient to ssh other linux machines by logIn.
# QuickStart
1. download this package
2. vim "login.go",edit configPath 
```
configPath = "/home/zzjack/tools/config.json"
```
3. vim "config.json"
note: the default port is 22
4. compile, "go build login.go"
5. execute, "./login xx".xx is your machine name.As example,one of mine is "crawler-test".

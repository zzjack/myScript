## sz
package files and transfer files.
### Install
```
git clone https://github.com/zzjack/myScript
``` 
### Quick Start
1. edit config.json.
2. edit sz.go,set configPath,as shown
```
var configPath = "set entire path of config.json which editted in step 1"
```
3. Complie
```
go build sz.go
``` 
4. View the permission of complied file "sz",if there is not execute permission,do as shown:
```
chmod +x ./sz
```
5. launch
```
./sz transferredFile SettingKeyInConfig.json
```
### Example
1. edit config.json
```
{
  "user_conf":{
    "crawler-pro":{
      "user_name":"root",
      "passwd":"target passwd",
      "host" : "target home",
      "send_path":"/usr/local"
    }
  }
}

```
2. edit sz.go
$ whereis config.json
$ ~/tools/
```
var configPath = "~/tools/config.json"
```
3. compile
4. if file "sz" is green,the file can execute
5. launch
i want to transfer a package named "djangoFile" to target machine
```
sz djangoFile crawler-pro
```
as expected,you will see like that:
```
2017/11/07 11:42:17 read config of send path ==> /usr/local
2017/11/07 11:42:17 Executing Command ==> tar cvf djangoFile.tar.gz djangoFile/
2017/11/07 11:42:17 Executing Command ==> scp djangoFile.tar.gz root@********:/usr/local
2017/11/07 11:42:17 expected string ==> root@**********'s password:
 
djangoFile.tar.gz                             100%   10KB 171.4KB/s   00:00
```



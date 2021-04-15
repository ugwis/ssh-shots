# ssh-shots
![ssh-shots](https://user-images.githubusercontent.com/914815/114811101-72448a00-9de8-11eb-9a71-8da4a1452dd7.jpg)

ssh-shots sends config to routers / switches in a single binary.


# Usage
Login in with an Password
```
go run main.go -pass password user host startup.conf
```
Login in with an SSH Private Key
```
go run main.go -i ~/.ssh/id_rsa.pem user host startup.conf
```

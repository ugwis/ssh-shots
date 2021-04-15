# ssh-shots
ssh-shots sends config to routers / switches in a single binary.

# Usage
Login in with an Password
```
go run main.go -i ~/.ssh/id_rsa.pem user host startup-config
```
Login in with an SSH Private Key
```
go run main.go -i ~/.ssh/id_rsa.pem user host startup-config
```

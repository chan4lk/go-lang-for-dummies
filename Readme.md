# init mod
```shell
go mod init chan4lk/hello
go mod tidy
```

# compile for mac
```shell
GOOS=darwin GOARCH=amd64 go build -o Chapter_1-mac
```

# compile for windows
```shell
GOOS=windows GOARCH=amd64 go build -o Chapter_1-windows.exe
```

# compile for linux
```shell
GOOS=linux GOARCH=amd64 go build -o Chapter_1-linux
```
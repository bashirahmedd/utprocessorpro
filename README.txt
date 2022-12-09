golang environment on Ubuntu OS:

naji@E6400:~/go/bin$ go version
go version go1.17.2 linux/amd64

naji@E6400:~/go/bin$ ./dlv version
Delve Debugger
Version: 1.9.1


VS Code Extension:

Name: Go
Id: golang.go
Description: Rich Go language support for Visual Studio Code
Version: 0.35.2
Publisher: Go Team at Google


Package used
$go get github.com/spf13/viper      (is not installing)
$go get github.com/ilyakaznacheev/cleanenv 

Further help:
https://code.visualstudio.com/docs/editor/debugging

To add modul
$ go mod init utube-downloader       (initialize local new module)
$ go get github.com/ilyakaznacheev/cleanenv   (add package to the module)
$ go mod tidy                   (removes unused Package)
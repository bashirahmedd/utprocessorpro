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
$go get github.com/spf13/viper      
$go get github.com/kkdai/youtube/v2
$go get github.com/ilyakaznacheev/cleanenv  (used and removed, too simple)
$go get github.com/spf13/cobra/cobra

Further help:
https://code.visualstudio.com/docs/editor/debugging

To add module
$ go mod init utube-downloader       (initialize local new module)
$ go get github.com/spf13/viper      (add package to the module)
$ go mod tidy                        (removes unused Package)
$ go mod vendor                      (removes inconsistency )
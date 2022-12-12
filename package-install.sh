#!/bin/bash
windscribe connect
#sudo apt-get install gnutls-bin
git config --global http.sslVerify false
git config --global http.postBuffer 1048576000
git config --global core.longpaths true
git config --global --unset credential.helper
go get -u github.com/ilyakaznacheev/cleanenv
#go get -u github.com/spf13/viper
windscribe disconnect
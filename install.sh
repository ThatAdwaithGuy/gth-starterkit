#!/bin/bash
rm -rf .git
git init
go mod init github.com/$1
go get github.com/a-h/templ
go get github.com/a-h/templ
go get -u github.com/gin-gonic/gin

npm i

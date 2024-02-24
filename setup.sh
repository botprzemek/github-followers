#!/bin/bash

go mod init
go get github.com/joho/godotenv
go get github.com/urfave/cli/v2
go install
github-followers install
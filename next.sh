#!/bin/bash

N=`ls | sed -n "s/^.*prob//p" | sort -g | tail -1`
N=$((N + 1))

mkdir prob$N
cd prob$N
go mod init euler/prob$N
echo "replace euler/utils => ../utils" >> go.mod
cp ../prob3/prob3.go ./prob$N.go
go mod tidy
cd ..
go work use prob$N

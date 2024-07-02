#!/usr/bin/env bash

# 这里要谨慎填写，后面有一条rm -fr $TargetDir, 避免删除重要目录。
TargetDir="./output"
ProjectName="amis-demo"
ExecuteBin="demo"
rm -fr "$TargetDir"
mkdir -p "$TargetDir/$ProjectName"

GOOS=linux GOARCH=amd64  go build -a -o "$TargetDir/$ProjectName/$ExecuteBin" .
chmod +x "$TargetDir/$ProjectName/$ExecuteBin"
cp -r ./web "$TargetDir/$ProjectName"

cd $TargetDir/ && tar -cf "amis-demo.tar" $ProjectName
echo "build success, target file: $TargetDir/$ProjectName/amis-demo.tar"

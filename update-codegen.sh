#!/bin/bash

# 进入code-generator目录
cd vendor/github.com/kubernetes/code-generator

# 执行命令
./generate-groups.sh "all" \
./nfs-controller/pkg/client \
./nfs-controller/pkg/apis \
"nfscontroller:v1alphv1"
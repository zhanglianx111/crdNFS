#!/bin/bash

# 进入code-generator目录
cd vendor/github.com/kubernetes/code-generator

# 执行命令
./generate-groups.sh "all" \
github.com/zhanglianx111/nfs-controller/pkg/client \
github.com/zhanglianx111/nfs-controller/pkg/apis \
"nfscontroller:v1alpha1"
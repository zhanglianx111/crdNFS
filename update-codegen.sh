#!/bin/bash

vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/zhaglianx111/nfscontroller/pkg/client \
github.com/zhanglianx111/nfscontroller/pkg/apis \
nfscontroller:v1alpha1
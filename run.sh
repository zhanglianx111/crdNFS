#!/bin/sh

go build && ./nfscontroller -master https://10.0.90.45:6443 -kubeconfig ./kubeconfig 

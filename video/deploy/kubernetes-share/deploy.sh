#! /bin/bash

# this script is just for reference, please do not run it directly
# require helm

# Install ingress-nginx
helm install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.2.5

# Install jaeger
kubectl create -f jaeger.yaml

# Install knative
kubectl apply -f https://github.com/knative/operator/releases/download/knative-v1.7.0/operator.yaml
kubectl create -f knative.yaml

# Install minio
kubectl create -f minio.yaml

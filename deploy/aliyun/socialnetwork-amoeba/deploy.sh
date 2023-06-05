#! /bin/bash

# this script is just for reference, please do not run it directly

SERVICES=(post-storage user-timeline user social-graph home-timeline media url-shorten user-mention text compose-post)

# Install namespace
kubectl apply -f namespace.yaml

# Install service config
kubectl apply -f config.yaml

# Install vm services
for s in ${SERVICES[@]}; do
    kubectl apply -f vm-service/$s-service.yaml
done

# Install serverless services
for s in ${SERVICES[@]}; do
    kubectl apply -f serverless-service/$s-service.yaml
done

# Install gateway
kubectl apply -f gateway.yaml

# Install ServiceMonitor
kubectl apply -f stats.yaml

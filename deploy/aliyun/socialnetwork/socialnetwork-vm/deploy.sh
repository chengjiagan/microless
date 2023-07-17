#! /bin/bash

# this script is just for reference, please do not run it directly

# SERVICES=(post-storage user-timeline user social-graph home-timeline media url-shorten user-mention text compose-post)
SERVICES=(post-storage user-timeline home-timeline)

# Install namespace
kubectl apply -f namespace.yaml

# Install service config
kubectl apply -f config.yaml

# Install services
for s in ${SERVICES[@]}; do
    kubectl apply -f service/$s-service.yaml
done
for s in ${SERVICES[@]}; do
    kubectl apply -f hpa/$s-hpa.yaml
done

# Install gateway
kubectl apply -f gateway.yaml

# Install ServiceMonitor
kubectl apply -f stats.yaml

# Install default role
kubectl apply -f rbac.yaml

#! /bin/bash

# this script is just for reference, please do not run it directly

SERVICES=(post-storage user-timeline user social-graph home-timeline media url-shorten user-mention text compose-post)

# apply socialnetwork namespace
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

# # Wait for services
# sleep 30

# # Run test
# kubectl apply -f test.yaml

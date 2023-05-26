#! /bin/bash

# this script is just for reference, please do not run it directly

# Install socialnetwork namespace
kubectl apply -f namespace.yaml

# Install service config
kubectl apply -f config.yaml

SERVICES=(post-storage user-timeline user social-graph home-timeline media url-shorten user-mention text compose-post)

# Install serverless services
for s in ${SERVICES[@]}; do
    kubectl apply -f serverless-service/$s-service.yaml
done

# Install vm services
for s in ${SERVICES[@]}; do
    kubectl apply -f vm-service/$s-service.yaml
done

# Install services
for s in ${SERVICES[@]}; do
    kubectl apply -f service/$s-service.yaml
done

# Install restful gateway
kubectl apply -f gateway.yaml

# Install service monitor
kubectl apply -f stats.yaml

# Install HPA objects for vm services
for s in ${SERVICES[@]}; do
    kubectl apply -f vm-hpa/$s-hpa.yaml
done

# # Wait
# sleep 30

# # Run test
# kubectl apply -f test.yaml

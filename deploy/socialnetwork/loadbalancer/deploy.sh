#! /bin/bash

# this script is just for reference, please do not run it directly

SERVICES=(post-storage user-timeline user social-graph home-timeline media url-shorten user-mention text compose-post)

# Install socialnetwork namespace
kubectl apply -f namespace.yaml

# Install service config
kubectl apply -f config.yaml
for s in ${SERVICES[@]}; do
    kubectl apply -f lb-config/$s-config.yaml
done

# Install serverless services
for s in ${SERVICES[@]}; do
    kubectl apply -f serverless-service/$s-service.yaml
done
for s in ${SERVICES[@]}; do
    kubectl apply -f serverless-hpa/$s-hpa.yaml
done

# Install vm services
for s in ${SERVICES[@]}; do
    kubectl apply -f vm-service/$s-service.yaml
done
for s in ${SERVICES[@]}; do
    kubectl apply -f vm-hpa/$s-hpa.yaml
done

# Install restful gateway
kubectl apply -f gateway.yaml

# Install service monitor
for s in ${SERVICES[@]}; do
    kubectl apply -f service/$s-service.yaml
done
kubectl apply -f stats.yaml

# Install serverless autoscaler
kubectl apply -f serverless-autoscaler.yaml

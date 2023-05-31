#! /bin/bash

# this script is just for reference, please do not run it directly

# apply socialnetwork namespace
kubectl apply -f namespace.yaml

# Install service config
kubectl apply -f config.yaml
kubectl apply -f lb-config/ping-config.yaml
kubectl apply -f lb-config/pong-config.yaml

# Install serverless services
kubectl apply -f serverless-service/ping-service.yaml
kubectl apply -f serverless-service/pong-service.yaml
kubectl apply -f serverless-hpa/ping-hpa.yaml
kubectl apply -f serverless-hpa/pong-hpa.yaml

# Install vm services
kubectl apply -f vm-service/ping-service.yaml
kubectl apply -f vm-service/pong-service.yaml
kubectl apply -f vm-hpa/ping-hpa.yaml
kubectl apply -f vm-hpa/pong-hpa.yaml

# Install gateway
kubectl apply -f gateway.yaml

# Install ServiceMonitor
kubectl apply -f service/ping-service.yaml
kubectl apply -f service/pong-service.yaml
kubectl apply -f stats.yaml

# Install serverless autoscaler
kubectl apply -f serverless-autoscaler.yaml

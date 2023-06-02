#! /bin/bash

# this script is just for reference, please do not run it directly

# apply socialnetwork namespace
kubectl apply -f namespace.yaml

# Install service config
kubectl apply -f config.yaml

# Install services
# kubectl apply -f service/ping-service.yaml
# kubectl apply -f service/pong-service.yaml
kubectl apply -f knative/ping-service.yaml
kubectl apply -f knative/pong-service.yaml

# Install gateway
kubectl apply -f gateway.yaml

# Install ServiceMonitor
# kubectl apply -f stats.yaml

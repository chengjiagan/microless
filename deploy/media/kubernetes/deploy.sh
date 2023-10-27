#! /bin/bash

# this script is just for reference, please do not run it directly

SERVICES=(cast-info compose-review movie-info movie-review page plot rating review-storage user-review user)

# Create namespace
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

#! /bin/bash

# this script is just for reference, please do not run it directly

# Create socialnetwork namespace
kubectl create -f namespace.yaml

# Install vm-services
kubectl create -f vm-service/config.yaml
kubectl create -f vm-service/post-storage-service.yaml
kubectl create -f vm-service/user-timeline-service.yaml
kubectl create -f vm-service/user-service.yaml
kubectl create -f vm-service/social-graph-service.yaml
kubectl create -f vm-service/home-timeline-service.yaml
kubectl create -f vm-service/media-service.yaml
kubectl create -f vm-service/url-shorten-service.yaml
kubectl create -f vm-service/user-mention-service.yaml
kubectl create -f vm-service/text-service.yaml
kubectl create -f vm-service/compose-post-service.yaml
kubectl create -f vm-service/restful-gateway-service.yaml

# Install kn-services
kubectl create -f kn-service/config.yaml
kubectl create -f kn-service/post-storage-service.yaml
kubectl create -f kn-service/user-timeline-service.yaml
kubectl create -f kn-service/user-service.yaml
kubectl create -f kn-service/social-graph-service.yaml
kubectl create -f kn-service/home-timeline-service.yaml
kubectl create -f kn-service/media-service.yaml
kubectl create -f kn-service/url-shorten-service.yaml
kubectl create -f kn-service/user-mention-service.yaml
kubectl create -f kn-service/text-service.yaml
kubectl create -f kn-service/compose-post-service.yaml
kubectl create -f kn-service/restful-gateway-service.yaml

# Install ingress
kubectl create -f ingress.yaml

# Install HPA objects
kubectl create -f hpa/post-storage-hpa.yaml
kubectl create -f hpa/user-timeline-hpa.yaml
kubectl create -f hpa/user-hpa.yaml
kubectl create -f hpa/social-graph-hpa.yaml
kubectl create -f hpa/home-timeline-hpa.yaml
kubectl create -f hpa/media-hpa.yaml
kubectl create -f hpa/url-shorten-hpa.yaml
kubectl create -f hpa/user-mention-hpa.yaml
kubectl create -f hpa/text-hpa.yaml
kubectl create -f hpa/compose-post-hpa.yaml
kubectl create -f hpa/restful-gateway-hpa.yaml

# Install load balancer
kubectl create -f loadbalancer/post-storage-lb.yaml
kubectl create -f loadbalancer/user-timeline-lb.yaml
kubectl create -f loadbalancer/user-lb.yaml
kubectl create -f loadbalancer/social-graph-lb.yaml
kubectl create -f loadbalancer/home-timeline-lb.yaml
kubectl create -f loadbalancer/media-lb.yaml
kubectl create -f loadbalancer/url-shorten-lb.yaml
kubectl create -f loadbalancer/user-mention-lb.yaml
kubectl create -f loadbalancer/text-lb.yaml
kubectl create -f loadbalancer/compose-post-lb.yaml
kubectl create -f loadbalancer/restful-gateway-lb.yaml

# # Wait for vm-services
# sleep 30

# # Run test
# kubectl create -f test.yaml

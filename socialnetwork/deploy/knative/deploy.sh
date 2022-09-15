#! /bin/bash

# this script is just for reference, please do not run it directly
# require helm

# Install ingress-nginx
helm install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --version 4.2.5

# Install knative
kubectl apply -f https://github.com/knative/operator/releases/download/knative-v1.7.0/operator.yaml
kubectl create -f knative.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.7.1/serving-default-domain.yaml

# Create namespace
kubectl create -f namespace.yaml

# Install jaeger
kubectl create -f ../kubernetes/backend/jaeger.yaml
# Install redis
kubectl create -f ../kubernetes/backend/redis.yaml
# Install mongodb
kubectl create -f ../kubernetes/backend/mongodb.yaml
# Install memcached
kubectl create -f ../kubernetes/backend/post-storage-memcached.yaml
kubectl create -f ../kubernetes/backend/user-memcached.yaml
kubectl create -f ../kubernetes/backend/url-shorten-memcached.yaml

# Install service config
kubectl create -f config.yaml

# Install services
kubectl create -f service/post-storage-service.yaml
kubectl create -f service/user-timeline-service.yaml
kubectl create -f service/user-service.yaml
kubectl create -f service/social-graph-service.yaml
kubectl create -f service/home-timeline-service.yaml
kubectl create -f service/media-service.yaml
kubectl create -f service/url-shorten-service.yaml
kubectl create -f service/user-mention-service.yaml
kubectl create -f service/text-service.yaml
kubectl create -f service/compose-post-service.yaml

# Install restful gateway
kubectl create -f restful-gateway.yaml
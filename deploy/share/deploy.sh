#! /bin/bash

# this script is just for reference, please do not run it directly
# require helm

# Install ingress-nginx
# helm install ingress-nginx ingress-nginx \
#   --repo https://kubernetes.github.io/ingress-nginx \
#   --namespace ingress-nginx --create-namespace \
#   --version 4.2.5

# Install prometheus and grafana
# grafana user: admin password: prom-operator
helm install prometheus-stack kube-prometheus-stack \
  --repo https://prometheus-community.github.io/helm-charts \
  --namespace monitoring --create-namespace \
  --version 41.9.1

# Install jaeger
kubectl create -f jaeger.yaml

# Install knative
kubectl apply -f https://github.com/knative/operator/releases/download/knative-v1.7.0/operator.yaml
kubectl create -f knative.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.7.0/serving-default-domain.yaml

# Install mongodb
kubectl create -f mongodb.yaml
# Install prometheus-mongodb-exporter
helm install mongodb-exporter prometheus-mongodb-exporter \
    --repo https://prometheus-community.github.io/helm-charts \
    --namespace monitoring --create-namespace \
    --version 3.1.2 \
    --set serviceMonitor.enabled=true \
    --set serviceMonitor.additionalLabels.release=prometheus-stack \
    --set mongodb.uri="mongodb://mongodb.mongodb:27017"

# Install minio
kubectl create -f minio.yaml

# Install redis
kubectl create -f redis.yaml

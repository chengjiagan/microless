#! /bin/bash

# this script is just for reference, please do not run it directly
# require helm

# Install prometheus and grafana
# grafana user: admin password: prom-operator
helm install prometheus-stack kube-prometheus-stack \
  --repo https://prometheus-community.github.io/helm-charts \
  --namespace monitoring --create-namespace

helm install prometheus-adapter prometheus-adapter \
  --repo https://prometheus-community.github.io/helm-charts \
  --namespace monitoring \
  --values adapter.yaml

# Install jaeger
kubectl create -f jaeger.yaml

# Install knative
kubectl apply -f https://github.com/knative/operator/releases/download/knative-v1.10.1/operator.yaml
kubectl create -f knative.yaml
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.10.1/serving-default-domain.yaml

# Install mongodb
kubectl create -f mongodb.yaml

# Install minio
# kubectl create -f minio.yaml

# Install redis
kubectl create -f redis.yaml

# Install istio
# istioctl install --set meshConfig.defaultConfig.holdApplicationUntilProxyStarts=true -y

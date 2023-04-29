#! /bin/bash

# this script is just for reference, please do not run it directly

kubectl delete ns social-network-test
kubectl delete ns social-network
kubectl delete -f envoy.yaml

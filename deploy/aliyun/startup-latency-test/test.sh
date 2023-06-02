#! /bin/bash
kubectl scale deployment test --replicas 8
kubectl wait pods --for=condition=Ready -l app=test --timeout=-1s
# kubectl scale deployment test --replicas 1

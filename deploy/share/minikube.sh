#! /bin/sh

minikube start --nodes 6
minikube stop
# name         role            spec   label
# minikube     master node     6C-8G  type=system
# minikube-m02 vm pods         4C-8G  type=vm
# minikube-m03 vm pods         4C-8G  type=vm
# minikube-m04 vm pods         4C-8G  type=vm
# minikube-m05 vm pods         4C-8G  type=vm
# minikube-m06 serverless pods 8C-16G type=serverless

# set CPUs
sudo virsh setvcpus minikube 6 --maximum --config
sudo virsh setvcpus minikube 6 --config
for i in 2 3 4 5; do
    sudo virsh setvcpus minikube-m0$i 4 --maximum --config
    sudo virsh setvcpus minikube-m0$i 4 --config
done
sudo virsh setvcpus minikube-m06 8 --maximum --config
sudo virsh setvcpus minikube-m06 8 --config

# set memory
sudo virsh setmaxmem minikube 8GiB --config
sudo virsh setmem minikube 8GiB --config
for i in 2 3 4 5; do
    sudo virsh setmaxmem minikube-m0$i 8GiB --config
    sudo virsh setmem minikube-m0$i 8GiB --config
done
sudo virsh setmaxmem minikube-m06 16GiB --config
sudo virsh setmem minikube-m06 16GiB --config

minikube start

# label in kubernetes
kubectl label node minikube type=system
for i in 2 3 4 5; do
    kubectl label node minikube-m0$i type=vm
done
kubectl label node minikube-m06 type=serverless

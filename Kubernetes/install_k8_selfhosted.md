## Install Kubernetes to Server
### Requirement
- Docker

### Step 1
- sudo apt update && sudo apt upgrade -y
- sudo apt install docker.io

### Step 2
- `export MASTER_IP=master.server.ip.address_`
``Replace the master.server.ip.address with the IP address of your desired master node``

### Step 3
- $curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
- Add repo 

`` $cat < /etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF ``

### Step 4
- sudo apt-get update && sudo apt-get install -y --allow-unauthenticated kubelet kubeadm kubectl kubernetes-cni

### Step 5
- `$kubeadm init --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address $MASTER_IP`

then :
```
$mkdir -p $HOME/.kube
$sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### Step 6
- ``$kubectl create -f https://raw.githubusercontent.com/winse/docker-hadoop/master/kube-deploy/kubeadm/kube-flannel-rbac.yml --namespace=kube-system``
- ``$kubectl create -f https://raw.githubusercontent.com/winse/docker-hadoop/master/kube-deploy/kubeadm/kube-flannel.yml --namespace=kube-system``

### Step 7 (option)
- `root@node-1:~$ kubeadm join --token ee4fea.3c299a9db86b1f1c $MASTER_IP:6443`

### Step 8
- bob@master:~$ kubectl get nodes
- bob@master:~$ kubectl get pods --all-namespaces



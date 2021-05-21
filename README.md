# Kubernetes Example

Simple location encryption with WebUI for learning K8s.

### Prerequisites

[Docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/)

### Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
docker-compose up
```

### Minikube 

```bash
kubectl create deployment redis --image=redis
kubectl get all
kubectl create deployment coordinates --image=leningotru/coordinates
kubectl create deployment encrypter --image=ronniemh/location-encrypt
kubectl create deployment web --image=danidaniel6462/ksh_webui
kubectl create deployment worker --image=danidaniel6462/ksh_worker
kubectl expose deployment redis --port 6379
kubectl expose deployment coordinates --port 3000
kubectl expose deployment encrypter --port 8080
kubectl expose deploy/web --type=NodePort --port 80
minikube service web
```
### EKS

```bash
eksctl create cluster \
 --name meet \
 --version auto \
 --nodegroup-name linux-nodes \
 --zones us-east-1a,us-east-1b \
 --node-type t3.small \
 --nodes 1 \
 --nodes-min 1 \
 --nodes-max 3 \
 --asg-access
```

### Usage

Open a Web UI: http://localhost:3000

locationEncrypt: Golang microservice to encrypt location


### License
[MIT](https://choosealicense.com/licenses/mit/)
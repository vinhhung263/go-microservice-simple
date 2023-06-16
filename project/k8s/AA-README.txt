*** To deploy
cd project

kubectl apply -f k8s
kubectl apply -f k8s/broker.yml
kubectl get pods
kubectl get svc
kubectl get nodes
kubectl get deployments
kubectl logs broker-service-df8d875b7-d4qd5 <pod name>
kubectl delete deployments
kubectl delete svc

kubectl expose deployment broker-service --type=LoadBalancer --port=8080 --target-port=8080
minikube tunnel

*Scale svc : update file .yml => use apply -f cmd

minikube dashboard
minikube status
minikube start
minikube stop
minikube addons enable ingress
kubectl apply -f ingress.yml


minikube service list

*special case*
docker-compose -f postgres.yml up -d

check ping from minikube to postgres local
=> nc -vz host.minikube.internal <port postgres>
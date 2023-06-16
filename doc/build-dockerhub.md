## ******************* Build and push to docker hub *******************
cmd:
docker login
cd logger-service
docker build -f logger-service.dockerfile -t vinhhung263/logger-service:1.0.0 .
docker push vinhhung263/logger-service:1.0.0
cd ..
cd authentication-service
docker build -f authentication-service.dockerfile -t vinhhung263/authentication-service:1.0.0 .
docker push vinhhung263/authentication-service:1.0.0
cd ..
cd broker-service
docker build -f broker-service.dockerfile -t vinhhung263/broker-service:1.0.1 .
docker push vinhhung263/broker-service:1.0.1
cd ..
cd listener-service
docker build -f listener-service.dockerfile -t vinhhung263/listener-service:1.0.0 .
docker push vinhhung263/listener-service:1.0.0
cd ..
cd mail-service
docker build -f mail-service.dockerfile -t vinhhung263/mail-service:1.0.0 .
docker push vinhhung263/mail-service:1.0.0

cd front-end
docker build -f front-end.dockerfile -t vinhhung263/front-end:1.0.1 .
docker push vinhhung263/front-end:1.0.1

## ******************* Build docker swarm *******************
cd project
docker swarm init
docker stack deploy -c swarm.yml myapp
docker service ls

docker service scale myapp_listener-service=3
docker service scale myapp_listener-service=2

## ******************* Build another code version *******************
cd logger-service
docker build -f logger-service.dockerfile -t vinhhung263/logger-service:1.0.1 .
docker push vinhhung263/logger-service:1.0.1
cd project
docker service scale myapp_logger-service=2
docker service update --image vinhhung263/logger-service:1.0.1 myapp_logger-service

## ******************* Stop docker swarm *******************
docker service scale myapp_broker-service=0
docker stack rm myapp
docker swarm leave

## ******************* Config reverse proxy *******************
cd project
docker build -f caddy.dockerfile -t vinhhung263/micro-caddy:1.0.0 .
docker push vinhhung263/micro-caddy:1.0.0
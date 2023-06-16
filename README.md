In this course, we'll develop a number of small, self-contained, loosely coupled microservices that will will communicate with one another and a simple front-end application with a REST API, with RPC, over gRPC, and by sending and consuming messages using AMQP, the Advanced Message Queuing Protocol. The microservices we build will include the following functionality:

A Front End service, that just displays web pages;

An Authentication service, with a Postgres database;

A Logging service, with a MongoDB database;

A Listener service, which receives messages from RabbitMQ and acts upon them;

A Broker service, which is an optional single point of entry into the microservice cluster;

A Mail service, which takes a JSON payload, converts into a formatted email, and send it out.

All of these services will be written in Go, commonly referred to as Golang, a language which is particularly well suited to building distributed web applications.

We'll also learn how to deploy our distributed application to a Docker Swarm and Kubernetes, and how to scale up and down, as necessary, and to update individual microservices with little or no downtime.

![Architecture](https://github.com/vinhhung263/go-microservice-simple/assets/62415557/7c558a62-4474-4ebc-b85d-427e088259dd)

Source: https://www.udemy.com/course/working-with-microservices-in-go/

# golang-rabbitmq
This repository show you golang examples which interacts with RabbitMQ queues and messages.

# Technologies used in this project

- Golang
- RabbitMQ

# Setting up this project locally

> **Note:**
The fastest way to get this application up and running locally is using **Docker**, **Docker Compose** and **Golang**.  Be sure that you have at least **Docker 1.13.0**, **Docker Compose 1.11.2** and **Golang 1.8** installed on your machine.

1. Clone this repository
```shell
$ git clone https://github.com/godois/golang-rabbitmq.git
```
2. Setup the dependencies (RabbitMQ)

```shell
$ docker-compose -f docker-compose/dependencies.yml up
```

3. Start the application

Example 1

You need to run a script main.go passing the parameter "receive" to put a thread listening the queue. 

```shell
$ go run example-1/main.go receive
```

So, You need to run another thread to send a message to the queue, runnign a script main.go passing the parameter "send". 

```shell
$ go run example-1/main.go send
```
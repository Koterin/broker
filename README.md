# broker

RabbitMQ &amp; Kafka use sample

# How to run
## RabbitMQ

1. Run the local RabbitMQ (case of MacOS below):
```
brew install rabbitmq
rabbitmqctl status (ensure the port 5672 is being listened by AMQP)
rabbitmq-server start (in case if it is not running)
```

2. Run the Producer app:
```
cd rabbitmq/
go run .
```
It will wait for user input from console. Press `enter` when you want to send a message and print `exit` for graceful shutting down.

3. Run the Consumer app:
```
cd rabbitmq/consumer1
go run .

cd rabbitmq/consumer2
go run .
```
It will connect to the RabbitMQ queue and print all the messages awaiting; each consumer also imitates some work going on by sleeping for 10 seconds after receiving each message. The RabbitMQ acknowledgement is done manually by marking `d.Ack(false)` after every message processing.
Consumer1 also receives no more than 2 messages at once, Consumer2 - no more than 1

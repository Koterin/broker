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
It will end immidiately after sending the message to the RabbitMQ queue

3. Run the Consumer app:
```
cd rabbitmq/consumer
go run .
```
It will connect to the RabbitMQ queue and print all the messages awaiting

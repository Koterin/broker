module github.com/koterin/broker/rabbitmq

go 1.18

replace github.com/koterin/broker/rabbitmq/server => ./server

require github.com/koterin/broker/rabbitmq/server v0.0.0-00010101000000-000000000000

require github.com/rabbitmq/amqp091-go v1.3.4 // indirect

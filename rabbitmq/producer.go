package main

import (
    "log"

    amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
    log.SetPrefix("[LOG] ")
    log.SetFlags(3)

    log.Printf("Producer server started successfully")

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "notes", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    body := "r u up?? have you lost my messages????"
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            ContentType: "text/plain",
            Body:        []byte(body),
        })

    failOnError(err, "Failed to publish a message")
    log.Printf("Sent %s\n", body)
}

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

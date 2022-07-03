package main

import (
    "log"
    "bufio"
    "strings"
    "os"

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

    log.Println("Enter your message: ")

    for {
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Println("An error occured while reading input. Please try again", err)
            return
        }

        input = strings.TrimSuffix(input, "\n")

        if strings.Compare(input, "exit") == 0 {
            log.Println("Gracefully shutting down...")
            break
        }
        err = ch.Publish(
            "",     // exchange
            q.Name, // routing key
            false,  // mandatory
            false,  // immediate
            amqp.Publishing {
                ContentType: "text/plain",
                Body:        []byte(input),
            })

        failOnError(err, "Failed to publish a message")
    }
}

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

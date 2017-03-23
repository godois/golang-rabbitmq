package example2

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/godois/golang-rabbitmq/util"
	"github.com/streadway/amqp"
)

var loadedURL string

func init() {
	util.LoadConfig()
	loadedURL = util.C.GetString("rabbitmq.rabbitmq_url")
}

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]

		if arg == "publish" {
			publish()
		} else if arg == "receive" {
			receive()
		} else {
			log.Print("[Golang-RabbitMQ - Error] - You need to inform a parameter executing this action (publish/receive) ...")
		}
	} else {
		log.Print("[Golang-RabbitMQ - Error] - You need to inform a parameter executing this action (publish/receive) ...")
	}
}

func publish() {

	conn, err := amqp.Dial(loadedURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	body := "Hello Message Publishing"
	err = ch.Publish(
		"",      // exchange
		"hello", // routing key
		false,   // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func receive() {
	conn, err := amqp.Dial(loadedURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"hello", // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

package main

import (
	"fmt"
	"log"
	"os"
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

		if arg == "send" {
			send()
		} else if arg == "receive" {
			receive()
		} else {
			log.Print(util.C.GetString("messages.error_parameter"))
		}
	} else {
		log.Print(util.C.GetString("messages.error_parameter"))
	}
}

func failOnError1(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func send() {

	conn, err := amqp.Dial(loadedURL)
	failOnError(err, util.C.GetString("messages.error_connection"))
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, util.C.GetString("messages.error_channel"))
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, util.C.GetString("messages.error_queue"))

	body := "hello message"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, util.C.GetString("messages.error_message_publish"))

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func receive() {

	conn, err := amqp.Dial(loadedURL)
	failOnError(err, util.C.GetString("messages.error_connection"))
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, util.C.GetString("messages.error_channel"))
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, util.C.GetString("messages.error_queue"))

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, util.C.GetString("messages.error_register_consumer"))

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(util.C.GetString("messages.info_receiving_message"))
	<-forever
}

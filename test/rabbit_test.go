package test

import (
	"log"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func TestRunR(t *testing.T) {

	doRabbit()
}

func doRabbit() {
	conn, err := amqp.Dial("amqp://guest:guest@140.83.83.152:8090/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// err = ch.ExchangeDeclare(
	// 	"logs",   // name
	// 	"fanout", // type
	// 	true,     // durable
	// 	false,    // auto-deleted
	// 	false,    // internal
	// 	false,    // no-wait
	// 	nil,      // arguments
	// )
	// failOnError(err, "Failed to declare an exchange")

	// q, err := ch.QueueDeclare(
	// 	"lz.issue", // name
	// 	false,      // durable
	// 	false,      // delete when unused
	// 	true,       // exclusive
	// 	false,      // no-wait
	// 	nil,        // arguments
	// )
	// failOnError(err, "Failed to declare a queue")

	// err = ch.QueueBind(
	// 	q.Name, // queue name
	// 	"",     // routing key
	// 	"logs", // exchange
	// 	false,
	// 	nil)
	// failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		"lz.issue", // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

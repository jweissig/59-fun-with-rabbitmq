//
// code from https://www.rabbitmq.com/tutorials/tutorial-two-go.html
//

package main

import (
	"bufio"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// read in our words
	file, err := os.Open("./3k.txt")

	// file likely does not exist
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// loop over words
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		//fmt.Println(scanner.Text())

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(scanner.Text()),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", scanner.Text())

	}

	// deal with any errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

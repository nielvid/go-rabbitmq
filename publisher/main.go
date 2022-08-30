package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Exception(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
func main() {

	fmt.Println("go rabbitmq server")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Exception(err)
	defer conn.Close()
	fmt.Println("successfully connected to rabbitmq instance")

	ch, err := conn.Channel()
	Exception(err)
	defer ch.Close()

	q, err := ch.QueueDeclare("users", false, false, false, false, nil)
	Exception(err)
	fmt.Println(q)
	err = ch.Publish("", "users", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte("my first golang rabbitmg message")})
	Exception(err)
	fmt.Println("successfully published messaged to rabbitmg in golang")

}

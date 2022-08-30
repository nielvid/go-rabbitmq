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

	fmt.Println("go rabbitmq consumer server")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Exception(err)
	defer conn.Close()
	fmt.Println("successfully connected to rabbitmq instance")

	ch, err := conn.Channel()
	Exception(err)
	defer ch.Close()

	msgs, err := ch.Consume("users", "", true, false, false, false, nil)
	Exception(err)

	msg := <-msgs
	res := string(msg.Body)

	fmt.Println(res)

}

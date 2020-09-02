package main

import (
	"amqp-session/client"
	"amqp-session/types"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	session, err := client.NewSession("amqp://guest:guest@dell")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = session.NewChannel("default")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = session.NewConsume(types.ConsumeOption{
		ChannelID: "default",
		Queue:     "test",
		Consumer:  "consumer",
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Callback: func(d amqp.Delivery) {
			logrus.Info("Received a message:", string(d.Body))
		},
	})
	if err != nil {
		logrus.Fatalln(err)
	}
	select {}
}

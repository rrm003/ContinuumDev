package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "test"
)

type User struct {
	ID      string
	Name    string
	Details string
}

func main() {

	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}
	var ID, Name, Details string
	for {

		fmt.Print("ID: ")
		fmt.Scanln(&ID)
		fmt.Print("Name: ")
		fmt.Scanln(&Name)
		fmt.Print("Details: ")
		fmt.Scanln(&Details)
		u := User{ID, Name, Details}
		publish(u, producer)
	}
}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

	return prd, err
}

func publish(user User, producer sarama.SyncProducer) {
	jsonuser, _ := json.Marshal(map[string]string{
		"id":      user.ID,
		"name":    user.Name,
		"details": user.Details,
	})
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jsonuser),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}
	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}

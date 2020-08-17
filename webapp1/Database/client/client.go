package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Shopify/sarama"
)

var port string

func getData(addr string) {

	res, err := http.Get(addr + "/detail")
	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
	}
}

func getUser(addr string) {
	var userid string
	fmt.Println("Enter User ID : ")
	fmt.Scanln(&userid)
	res, err := http.Get(addr + "/detail/" + userid)
	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}

type user struct {
	id       string
	name     string
	details  string
	comments map[string]string
}

func createUser(addr string, client sarama.Client) {
	newuser := user{}
	var user, comment string
	fmt.Print("Enter Id : ")
	fmt.Scanln(&newuser.id)
	fmt.Print("Enter Name : ")
	fmt.Scanln(&newuser.name)
	fmt.Print("Enter Detail : ")
	fmt.Scanln(&newuser.details)
	x := "y"
	fmt.Print("Add Comment Y/N? ")
	fmt.Scanln(&x)
	newuser.comments = make(map[string]string)
	for x == "y" || x == "Y" {
		fmt.Print("Enter Name : ")
		fmt.Scanln(&user)
		fmt.Print("Enter Comment : ")
		fmt.Scanln(&comment)
		newuser.comments[user] = comment

		fmt.Print("Add Comment Y/N? ")
		fmt.Scanln(&x)
	}

	newuserjson, err := json.Marshal(map[string]interface{}{
		"id":       newuser.id,
		"name":     newuser.name,
		"details":  newuser.details,
		"comments": newuser.comments,
	})
	if err != nil {
		fmt.Println("Marshaling err")
	}
	fmt.Println(newuserjson)
	producer, err := sarama.NewSyncProducerFromClient(client)
	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder(newuserjson),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", "test", partition, offset)
}

func updateUser(addr string) {
	newuser := user{}
	fmt.Print("Enter Id : ")
	fmt.Scanln(&newuser.id)
	fmt.Print("Enter Name : ")
	fmt.Scanln(&newuser.name)
	fmt.Print("Enter Detail : ")
	fmt.Scanln(&newuser.details)

	newuserjson, err := json.Marshal(map[string]string{
		"id":      newuser.id,
		"name":    newuser.name,
		"details": newuser.details,
	})
	if err != nil {
		fmt.Println("Marshaling err")
	}
	fmt.Println(newuserjson)
	req, err := http.NewRequest(http.MethodPut, addr+"/detail/update", bytes.NewBuffer(newuserjson))
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	fmt.Println(res)
}

func deleteUser(addr string) {
	var userid string
	fmt.Println("Enter User ID : ")
	fmt.Scanln(&userid)
	id, err := json.Marshal(map[string]string{
		"id": userid,
	})
	req, err := http.NewRequest(http.MethodDelete, addr+"/detail/delete/"+userid, bytes.NewBuffer(id))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func main() {

	config := sarama.NewConfig()
	brokers := []string{"localhost:9092"}
	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter Port Number : ")
	fmt.Scanln(&port)
	address := "http://localhost:" + port

	x := "y"
	for x == "y" || x == "Y" {
		var opt int
		fmt.Println("Option : 1.All_User 2.User 3.Create 4.Update 5.delete :")
		fmt.Scanf("%d", &opt)
		switch opt {
		case 1:
			getData(address)
		case 2:
			getUser(address)
		case 3:
			createUser(address, client)
		case 4:
			updateUser(address)
		case 5:
			deleteUser(address)
		default:
			fmt.Println("Invalid option!")
		}
		fmt.Print("Continue Y/N? ")
		fmt.Scanln(&x)
	}
}

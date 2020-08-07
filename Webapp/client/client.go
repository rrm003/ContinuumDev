package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	id      string
	name    string
	details string
}

func createUser(addr string) {
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
	res, err := http.Post(addr+"/detail/add", "application/json", bytes.NewBuffer(newuserjson))
	fmt.Println(res)
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
			createUser(address)
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

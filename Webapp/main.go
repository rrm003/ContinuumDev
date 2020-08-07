package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var filename = "data.json"

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
}

func readData() []User {
	filedata, _ := ioutil.ReadFile(filename)
	var u []User
	err := json.Unmarshal(filedata, &u)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return u
}

func writeData(u User) {
	fmt.Print(u)
	filedata := readData()
	filedata = append(filedata, u)
	data, _ := json.MarshalIndent(filedata, "", "\t")
	_ = ioutil.WriteFile(filename, data, 0600)
}
func updateFile(u []User) {
	data, _ := json.MarshalIndent(u, "", "\t")
	_ = ioutil.WriteFile(filename, data, 0600)
}

func detailsHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data := readData()
	json.NewEncoder(res).Encode(data)
}

func addUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	for _, getuser := range readData() {
		if user.ID == getuser.ID {
			json.NewEncoder(res).Encode("User Exists !")
			return
		}
	}
	writeData(user)
	json.NewEncoder(res).Encode(readData())
}

func updateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	data := readData()
	for index, getuser := range data {
		if user.ID == getuser.ID {
			data[index].Name = user.Name
			data[index].Details = user.Details
			updateFile(data)
			json.NewEncoder(res).Encode(readData())
			return
		}
	}
	json.NewEncoder(res).Encode("Update : No Matching data!")
}

func deleteUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	updateData := readData()
	for index, data := range updateData {
		if data.ID == params["id"] {
			updateData = append(updateData[:index], updateData[index+1:]...)
			updateFile(updateData)
			return
		}
	}
	json.NewEncoder(res).Encode("Data not Found")
}
func getUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, data := range readData() {
		if data.ID == params["id"] {
			json.NewEncoder(res).Encode(data)
			return
		}
	}
	json.NewEncoder(res).Encode("Data not Found")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/detail", detailsHandler).Methods("GET")
	r.HandleFunc("/detail/{id}", getUser).Methods("GET")
	r.HandleFunc("/detail/add", addUser).Methods("POST")
	r.HandleFunc("/detail/update", updateUser).Methods("PUT")
	r.HandleFunc("/detail/delete/{id}", deleteUser).Methods("DELETE")
	http.ListenAndServe(":8800", r)
}

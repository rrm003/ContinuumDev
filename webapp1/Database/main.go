package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Details  string            `json:"details"`
	Comments map[string]string `json:"comments"`
}

func detailsHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data := DbAllUser()
	json.NewEncoder(res).Encode(data)
}

func addUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user User
	fmt.Println(req.Body)
	_ = json.NewDecoder(req.Body).Decode(&user)

	for _, getuser := range DbAllUser() {
		if user.ID == getuser.ID {
			json.NewEncoder(res).Encode("User Exists !")
			return
		}
	}
	DbAddUser(user)
	json.NewEncoder(res).Encode(DbAllUser())
}

func updateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	data := DbAllUser()
	for index, getuser := range data {
		if user.ID == getuser.ID {
			data[index].Name = user.Name
			data[index].Details = user.Details
			data[index].Comments = user.Comments
			DbUpdate(user)
			json.NewEncoder(res).Encode(DbAllUser())
			return
		}
	}
	json.NewEncoder(res).Encode("Updata No matching data")
}

func deleteUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var id string
	id = params["id"]
	DbDeleteUser(id)
}
func getUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var id string
	id = params["id"]
	user := DbGetUser(id)
	json.NewEncoder(res).Encode(user)
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

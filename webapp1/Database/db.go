package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "webapp1"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Established")
}

func DbAddUser(u User) {
	fmt.Printf("Creating User : %v\n", u)
	fmt.Println(u.Comments)
	err := Session.Query("insert into user (userid, name, details, comments) values(?,?,?,?)", u.ID, u.Name, u.Details, u.Comments).Exec()
	if err != nil {
		panic(err)
	}
}
func DbAllUser() []User {
	fmt.Println("All User")
	var users []User
	m := map[string]interface{}{}
	iter := Session.Query("select * from user").Iter()

	for iter.MapScan(m) {
		users = append(users, User{
			ID:       m["userid"].(string),
			Name:     m["name"].(string),
			Details:  m["details"].(string),
			Comments: m["comments"].(map[string]string),
		})
		m = map[string]interface{}{}
	}
	return users
}

func DbGetUser(id string) User {
	fmt.Println(" User ", id)
	var users []User
	m := map[string]interface{}{}
	iter := Session.Query("select * from user where userid=?", id).Iter()

	for iter.MapScan(m) {
		users = append(users, User{
			ID:       m["userid"].(string),
			Name:     m["name"].(string),
			Details:  m["details"].(string),
			Comments: m["comments"].(map[string]string),
		})
		m = map[string]interface{}{}
	}
	return users[0]
}

func DbUpdate(u User) {
	fmt.Println("Updating user")
	err := Session.Query("update user set name=?, details=?, comments=? where userid=?", u.Name, u.Details, u.Comments, u.ID).Exec()
	if err != nil {
		panic(err)
	}
}

func DbDeleteUser(id string) {
	fmt.Println("Delete user")
	err := Session.Query("delete from user where userid = ?", id).Exec()
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	user := User{id: "1", name: "rrm", details: "003"}
// 	DbAllUser()
// 	DbAddUser(user)
// 	DbAllUser()
// 	DbGetUser("1")
// 	DbAllUser()
// 	DbUpdate(User{id: "1", name: "rrm", details: "SoftwareEngineer"})
// 	DbAllUser()
// 	DbDelete("1")
// 	DbAllUser()
// }

package main

import (
	"log"
	"os"
	"time"
)

var DataBase = []User{
	{Email: "test_1@gmail.com", Name: "John Doe 1"},
	{Email: "test_2@gmail.com", Name: "John Doe 2"},
	{Email: "test_3@gmail.com", Name: "John Doe 3"},
	{Email: "test_4@gmail.com", Name: "John Doe 4"},
	{Email: "test_5@gmail.com", Name: "John Doe 5"},
	{Email: "test_6@gmail.com", Name: "John Doe 6"},
	{Email: "test_7@gmail.com", Name: "John Doe 7"},
	{Email: "test_8@gmail.com", Name: "John Doe 8"},
	{Email: "test_9@gmail.com", Name: "John Doe 9"},
	{Email: "test_10@gmail.com", Name: "John Doe 10"},
	{Email: "test_11@gmail.com", Name: "John Doe 11"},
	{Email: "test_12@gmail.com", Name: "John Doe 12"},
	{Email: "test_13@gmail.com", Name: "John Doe 13"},
	{Email: "test_14@gmail.com", Name: "John Doe 14"},
}

type User struct {
	Email string
	Name  string
}

type Worker struct {
	users []User
	ch    chan *User
}

func GetAllWorkers(users []User, ch chan *User) *Worker {
	return &Worker{users: users, ch: ch}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if email == user.Email {
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]
	ch := make(chan *User)

	w := GetAllWorkers(DataBase, ch)

	log.Printf("Looking for %s", email)
	go w.Find(email)

	select {
	case user := <-ch:
		log.Printf("The email %s is owned by %s", email, user.Name)
	case <-time.After(1 * time.Second):
		log.Printf("The email %s was not found", email)
	}

}

package main

import (
	"log"
	"os"
)

type User struct {
	Email string
	Name  string
}

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

type Worker struct {
	users []User
}

func GetAllWorkers(users []User) *Worker {
	return &Worker{users: users}
}

func (w *Worker) Find(email string) *User {
	for i := range w.users {
		user := &w.users[i]
		if user.Email == email {
			return user
		}
	}
	return nil
}

func main() {
	email := os.Args[1]
	w := GetAllWorkers(DataBase)

	log.Printf("Looking for %s", email)
	user := w.Find(email)
	if user != nil {
		log.Printf("The email %s is owned by %s", email, user.Name)
	} else {
		log.Printf("The email %s was not found", email)
	}

}

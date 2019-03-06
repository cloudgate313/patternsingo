// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type User struct {
// 	Name, Role, Email string
// 	Age               int
// }

// func (u User) Salary() (int, error) {
// 	if u.Role == "" {
// 		return 0, errors.New("Cannot handle an empty role")
// 	}
// 	switch u.Role {
// 	case "Architect":
// 		return 250, nil
// 	case "Engineer":
// 		return 125, nil
// 	default:
// 		return 0, errors.New(
// 			fmt.Sprintf("I'm not able to handle the '%s", u.Role),
// 		)
// 	}
// }

// func (u *User) updateEmail(email string) {
// 	u.Email = email
// }

// func main() {
// 	user := User{Name: "Evan", Role: "Designer", Age: 39}

// 	if salary, err := user.Salary(); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Salary: ", salary)
// 	}

// 	fmt.Println(user)
// }

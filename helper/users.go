package helper

import "fmt"

type User struct{}

func (u User) GetUser() {
	fmt.Println("Hello Normal User")
}

package helper

import "fmt"

type Admin struct{}

func (a Admin) GetUser() {
	fmt.Println("Hello Admin")
}

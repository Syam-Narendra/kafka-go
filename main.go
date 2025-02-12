package main

import (
	"fmt"
	"time"

	"github.com/bxcodec/faker/v4"
)

func main() {

	for {
		time.Sleep(500 * time.Millisecond)

		user := map[string]string{
			"name":     faker.FirstName(),
			"email":    faker.Email(),
			"password": faker.Password(),
		}

		fmt.Println(user["name"])
	}

}

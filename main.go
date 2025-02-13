package main

import (
	"fmt"
	"time"
	"v1/producer"
)

func main() {

	for {
		time.Sleep(500 * time.Millisecond)
		event := producer.GenerateRandUserEvent()
		fmt.Println(event)

	}

}

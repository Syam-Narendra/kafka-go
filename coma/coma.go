package coma

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

const OuterVarialbe = "Outer Variable"

type smallUser struct {
	name string
}

type BigUser struct {
	name string
	Age  int
}

func Main() int {
	name, err := os.LookupEnv("name")
	fmt.Println("err in env", err)
	fmt.Println("Hello from coma")
	fmt.Println("env in coma", name)

	var numType = strconv.FormatUint(uint64(378), 10)
	var Num = 378
	fmt.Println(reflect.TypeOf(numType))
	return Num

}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Zero values
	totalPrice := 1000

	fmt.Println(totalPrice, reflect.TypeOf(totalPrice))
}

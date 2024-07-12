package main

import (
	"fmt"

	"github.com/monkig/arrgo/arrgo"
)

func main() {
	myArrgo := arrgo.New[int]()
	fmt.Println(myArrgo)
}

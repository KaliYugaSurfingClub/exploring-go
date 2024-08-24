package first

import "fmt"

func iotaEx() {
	//true(
	fmt.Println(Summer1 == Apples1)
	//error)
	//fmt.Println(Summer == Apples)

	fmt.Println(Summer == Autumn)
}

const (
	Summer1 int8 = iota
	Autumn1
	Winter1
	Spring1
)

const (
	Apples1 int8 = iota
	Oranges1
)

type Season int8

const (
	Summer Season = iota
	Autumn
	Winter
	Spring
)

type Fruit int8

const (
	Apples Fruit = iota
	Oranges
)

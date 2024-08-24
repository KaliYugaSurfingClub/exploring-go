package second

import "fmt"

func Overflow() {
	var ui uint = 10
	var deductible uint

	fmt.Scanf("%d", &deductible)

	ui = ui - deductible

	//ui = uint::max - 10
	fmt.Println(ui)
}

func Div() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ui uint = 10
	var divider uint

	fmt.Scanf("%d", &divider)

	//panic
	ui = ui / divider

	fmt.Println(ui)
}

func CallDiv() {
	Div()
	fmt.Println("was div")
}

func DivWithoutRecovery() {
	var ui uint = 10
	var divider uint

	fmt.Scanf("%d", &divider)

	//panic
	ui = ui / divider

	fmt.Println(ui)
}

func CallDivWithoutRecovery() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("was div and recover in Call func")
		}
	}()

	DivWithoutRecovery()
	fmt.Println("never happened")
}

//panic it is really exception LOOL, but func stops after a panic and goes to the defers

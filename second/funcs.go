package second

import (
	"fmt"
	"slices"
)

const Var = 10

func F(args ...int) int {
	res := 0
	for _, v := range args {
		res += v
	}
	return res
}

func CallF() {
	fmt.Println(Var)
	arr := []int{1, 2, 3}
	_ = F(arr...)
	//cannot
	//_ = f(arr)
}

func InitResInDef() (res int) {
	slices.Clone([]int{1, 2, 3})
	res = 10
	return
}

func init() {
	fmt.Println("init second")
}

func UseDefers() {
	//defers will be called from the last to the first
	for i := 0; i < Var; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	//done
	//9
	//8
	//...
	//1
}

func getCounter() func() int {
	i := 10

	return func() int {
		i++
		return i
	}
}

func UseCounters() {
	c1 := getCounter()
	c2 := getCounter()

	c1()
	c1()
	c1()

	c2()

	fmt.Println(c1(), c2())
}

//type VeryLargeStruct struct {
//	i int
//}
//
//func getVeryLargeStruct() VeryLargeStruct {
//	return VeryLargeStruct{i: 1}
//}
//
//func getVeryLargeStructPtr() *VeryLargeStruct {
//	return &VeryLargeStruct{i: 1}
//}
//
//func initVeryLargeStruct(s *VeryLargeStruct) {
//	s.i = 1
//}
//
//func useVeryLargeStruct() {
//	s := getVeryLargeStruct()
//	sPtr := getVeryLargeStructPtr()
//
//	var sInitedLikeC VeryLargeStruct
//
//}

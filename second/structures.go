package second

import "fmt"

type OwnSlice struct {
	i     int
	slice []int
}

func UseCopyStruct() {
	s := OwnSlice{1, []int{1, 2, 3}}
	s1 := s

	s.i = 0
	s.slice[0] = 99

	//have diff i and same slice
	fmt.Println(s, s1)
}

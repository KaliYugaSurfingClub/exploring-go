package first

import (
	"fmt"
	"regexp"
	"slices"
)

func Use_append() {
	//append does copy
	{
		arr := []int{1, 2, 3, 4}
		arr1 := append(arr, 1)

		arr[0] = 88
		//diff
		fmt.Println(arr, arr1)
	}

	//append with one arg does not return copy, returns ref
	{
		arr := []int{1, 2, 3, 4}
		arr1 := append(arr)

		arr[0] = 88
		//same
		fmt.Println(arr, arr1)
	}

	//if second arg of append is empty, append returns ref
	{
		arr := []int{1, 2, 3, 4}
		arr1 := append(arr, []int{}...)

		arr[0] = 10
		//same
		fmt.Println(arr, arr1)
	}

	//behaviour when cap does not change (returns ref)
	{
		arr := make([]int, 5, 10)
		arr1 := append(arr, 1)

		arr[0] = 88
		//same
		fmt.Println(arr, arr1)
	}

	//returns a copy only if the source array needs to be reallocated to add new elements
}

func Use_make() {
	//make create an array required len
	{
		arr := make([]int, 5)
		fmt.Println(arr)

		arr1 := make([]int, 5)
		arr1 = append(arr1, 10)

		fmt.Println(arr1)
	}

	//pass cap
	{
		arr := make([]int, 5, 10)
		arr = append(arr, 10)

		fmt.Println(arr)
		fmt.Println(len(arr), cap(arr))
	}
}

func Use_copy() {
	//default usage
	{
		src := []int{1, 2, 3, 4}
		dest := make([]int, len(src))

		copy(dest, src)
		src[0] = 88

		fmt.Println(src, dest)
	}

	//copy can not change size of slice
	{
		src := []int{1, 2, 3, 4}
		dest := make([]int, 2, 1000)

		copy(dest, src)
		src[0] = 88

		//do no increase
		fmt.Println(src, dest)
	}
	{
		src := []int{1, 2}
		dest := make([]int, 4)

		copy(dest, src)

		//do no decrease
		fmt.Println(src, dest)
	}

	//there is no way to change size of slice, and I only can do this shit
	{
		src := []int{1, 2, 3, 4}
		dest := make([]int, 2, 10000)

		dest = append(dest, make([]int, len(src)-len(dest))...)
		copy(dest, src)

		fmt.Println(src, dest)
	}

	//I have found way!!!!
	{
		src := []int{1, 2, 3, 4}
		dest := []int{1, 2}

		dest = slices.Clone(src)

		src[0] = 88
		fmt.Println(src, dest)
	}
}

func Use_slicing() {
	//slicing returns copy
	{
		arr := []int{1, 2, 3, 4}
		arr1 := arr[1:3]

		arr[1] = 88
		fmt.Println(arr, arr1)
	}

	//do not copy like python
	{
		arr := []int{1, 2, 3, 4}
		arr1 := arr[:]

		arr[0] = 88
		fmt.Println(arr, arr1)
	}

	//boring section
	{
		arr := []int{1, 2, 3, 4}
		arr = append(arr, arr[1:3]...)
		fmt.Println(arr)
	}
}

// BLOODY HELL
func FindDigits() []byte {
	digitRegexp := regexp.MustCompile("[0-9]+")
	VeryLargeString := []byte("abc123..1000000000000...")
	//will first group of consecutive numeric digits, returning them as a new slice.
	//and return as a ref
	//VeryLargeString will not remove becaUse there is a slice referring to VeryLargeString
	return digitRegexp.Find(VeryLargeString)
}

// better
func FindDigits2() []byte {
	digitRegexp := regexp.MustCompile("[0-9]+")
	VeryLargeString := []byte("abc123..1000000000000...")

	found := digitRegexp.Find(VeryLargeString)
	res := make([]byte, len(found))

	copy(res, found)

	return res
}

package first

import (
	"fmt"
)

func use_string() {
	//images
	{
		str := "Саша"
		//copy of str
		runes := []rune(str)
		//copy of str
		bytes := []byte(str)

		fmt.Println(str, len(str))
		fmt.Println(runes, len(runes))
		fmt.Println(bytes, len(bytes))

		runes[0] = '8'

		fmt.Println(str, len(str))
		fmt.Println(runes, len(runes))
		fmt.Println(bytes, len(bytes))
	}

	//copes and refs runes
	{
		str := []rune("Саша")
		str1 := str

		str[0] = '8'

		//same cuz these are slices
		fmt.Println(str, str1)
	}

	//copy and refs runes and string
	{
		str := "Саша"
		str1 := []rune(str)

		str1[0] = '8'

		//diff
		fmt.Println(str, string(str1))
	}

	//with slice
	{
		str := "Саша"
		str1 := []rune(str[:])

		str1[0] = '8'

		//diff
		fmt.Println(str, str1)
	}

	//mut
	{
		name := "Саша"
		surname := "Леонов"

		_ = name + " " + surname
		//str += '1';
		//str = append([]rune(str), '1');

		fmt.Println(len(name))
	}
}

func use_slicing_string() {
	str := "Саша"
	//use indexes in []byte not in []rune
	substr := str[2:6]

	//substr is a string_view
	fmt.Println(&str, &substr)
	fmt.Println(str, substr)

	//right way but
	substr1 := string([]rune(str)[1:3])
	fmt.Println(&str, &substr1)
	fmt.Println(str, substr1)
}

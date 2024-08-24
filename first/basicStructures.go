package first

import "fmt"

type User struct {
	name string
	age  int
}

func use_swith() {
	user := User{name: "Sasha", age: 19}

	//choose only one without break
	switch user.name {
	case "Misha":
		fmt.Println("User is Misha")
	case "Sasha":
		fmt.Println("User is sasha")
	case "Andrey":
		fmt.Println("User is Andrey")
	default:
		fmt.Println("User is unknown")
	}

	//instead if else ... if else etc.
	//fallthrough example
	switch {
	case user.name == "Misha":
		fmt.Println("User is Misha")
	case user.name == "Andrey":
		fmt.Println("User is Andrey")
	case user.name == "Sasha":
		fmt.Println("User is Sasha")
		fallthrough
	case user.age == 19:
		//will be reached thanks to fallthrough
		fmt.Println("User is 19")
	}

	//break example
	switch user.name {
	case "Misha", "Sasha":
		fmt.Println("User is Misha or Sasha")
		if user.age == 19 {
			break
		}
		//will never be happend
		fmt.Println("He is old")
	}

	//lables
	fmt.Println("switch with break")
Loop:
	for _, age := range []int{1, 2, 19, 88, 19} {
		switch {
		case age == user.age:
			fmt.Printf("User is %d\n", age)
			break Loop
		case age < user.age:
			fmt.Printf("User is older than %d\n", age)
		case age > user.age:
			fmt.Printf("User is older than %d\n", age)
		}
	}
}

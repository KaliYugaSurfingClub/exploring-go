package thrid

import "fmt"

//why close

func SelectInOneGo() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	for i := 0; i < 3; i++ {
		ch1 <- i
		ch2 <- -i
	}

	close(ch1)
	close(ch2)

	for i := 0; i < 10; i++ {
		//channels are closed and have data
		//behaviour is undefined what chan will be chosen
		//output -1 0 2 1 -2 0 0 0 ... 0
		select {
		case x := <-ch1:
			fmt.Println(x)
		case x := <-ch2:
			fmt.Println(x)
		default:
			//will never be executed cuz channels are closed
			//and have 5 values 3 mine and 2 empty
			fmt.Println("NO DATA")
		}
	}
}

func SelectInOneGoIgnoreEmptyValues() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	for i := 0; i < 3; i++ {
		ch1 <- i
		ch2 <- -i
	}

	close(ch1)
	close(ch2)

	for i := 0; i < 10; i++ {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println(x)
			}
		case x, ok := <-ch2:
			if ok {
				fmt.Println(x)
			}
		default:
			fmt.Println("NO DATA")
		}
	}
}

func SelectInOneGoNoClosing() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	for i := 0; i < 3; i++ {
		ch1 <- i
		ch2 <- -i
	}

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case x := <-ch2:
			fmt.Println(x)
		default:
			//case if any chan has no element
			fmt.Println("NO DATA")
		}
	}
}

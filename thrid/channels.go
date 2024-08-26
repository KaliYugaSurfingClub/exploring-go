package thrid

import (
	"fmt"
)

func GenAndUse() {
	buff := 10

	in := make(chan int)
	out := make(chan int)

	go func(in chan<- int) {
		for i := 0; i < buff; i++ {
			fmt.Printf("write %d in In\n", i)
			in <- i
		}
		close(in)
	}(in)

	go func(in <-chan int, out chan<- int) {
		for i := range in {
			fmt.Printf("write %d in Out\n", i*i)
			out <- i * i
		}
		close(out)
	}(in, out)

	for i := range out {
		fmt.Println(i)
	}
}

func Chain() {
	first := make(chan int)
	prev := first

	for i := 0; i < 1e6; i++ {
		next := make(chan int)

		go func(prev <-chan int) {
			number := <-prev
			next <- number + 1
		}(prev)

		prev = next
	}

	//создалось 1e6 goroutines и они все ждут когда поток, который они приняли, отдаст значение
	//создалось 1e6 каналов, которые ждут, когда в них кто то что то запишет

	//записываем в первый канал
	//первая горутина записывает значение в некст
	//вторая горутина которая ждет когда в некст(для этой горутины это прев)
	//что-то появиться записывает в слудующий некст
	//итд
	//prev уже равен последне созданному каналу читаем оттуда
	first <- 1
	fmt.Println(<-prev)
}

func Deadlock() {
	{
		ch := make(chan int)
		//writing blocks main goroutine and there is no one else goroutine => deadlock
		ch <- 10
		fmt.Println(<-ch)
	}
	{
		ch := make(chan int)
		//writing blocks main goroutine and there is no one else goroutine => deadlock
		ch <- 10
		//same deadlock
		close(ch)
		fmt.Println(<-ch)
	}
}

func FixDeadlock() {
	{
		ch := make(chan int)
		//there are two goroutines one write other read => no deadlock
		go func() {
			ch <- 10
		}()
		fmt.Println(<-ch)
	}
	{
		ch := make(chan int, 1)
		//goroutine can write in buffer and goes doing other work
		//main goroutine wrote in ch and went to read from ch
		ch <- 10
		fmt.Println(<-ch)
	}
}

func DeadlockWithBuff() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20

	//deadlock cuz main goroutine wrote and started wait (blocked)
	//when other goroutine read from ch at least one value to write another int
	//there is no other goroutines => deadlock
	ch <- 230

	//will not execute
	close(ch)

	fmt.Println(<-ch)
}

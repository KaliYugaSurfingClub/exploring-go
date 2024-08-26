package thrid

import (
	"fmt"
	"time"
)

func ForWithRace() {
	//correct in go 1.23 incorrect in go 1.18
	//когда до горутины доходит управление i уже может добежать до 10
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(2 * time.Second)
}

func RightFor() {
	//good for all versions
	{
		for i := 0; i < 10; i++ {
			go func(i int) {
				fmt.Println(i)
			}(i)
		}
		time.Sleep(2 * time.Second)
	}
	//other way
	{
		for i := 0; i < 10; i++ {
			i := i
			go func() {
				fmt.Println(i)
			}()
		}
		time.Sleep(2 * time.Second)
	}
}

func Race2() {
	//concurrent map writes
	mp := make(map[int]int)

	for i := 0; i < 1000; i++ {
		go func() {
			mp[1] = 1
		}()
	}

	time.Sleep(5 * time.Second)
}

func Race3() {
	//concurrent map read and write
	mp := make(map[int]int)

	go func() {
		for i := 0; i < 10000; i++ {
			//mp[0] = 0 it works because there is no recording to map
			//go use optimization for loops
			mp[0] = i
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			if mp[0] < 0 {

			}
		}
	}()

	time.Sleep(5 * time.Second)
}

func Race4() {
	//there is a race condition counter += 10 is a read and write
	//no fatal errors, but we cannot be sure that counter will be 1000 * 10
	//no fatal errors with -race
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			if counter < 10 {

			}
			counter = counter + 10
			fmt.Println(i, counter)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(counter)
}

func NotRace5() {
	//there is no race cuz all goroutine works with its own value in memory
	//they do not mut size, cap and ptr of arr
	arr := make([]int, 1e5)

	fmt.Println(arr[100:105])

	for i, _ := range arr {
		go func(i int) {
			arr[i] = i * i
		}(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println(arr[100:105])
}

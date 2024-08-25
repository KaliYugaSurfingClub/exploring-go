package thrid

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type pair struct {
	time time.Time
	i    string
}

func MyExample() {
	mp := make(map[string]time.Time)
	ch := make(chan pair)
	wg := sync.WaitGroup{}

	for i := 0; i < 88; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- pair{time.Now(), strconv.Itoa(i) + "Str"}
		}(i)
	}

	//deadlock cuz main waits and other also wait to write
	//for p := range ch {
	//	mp[p.i] = p.time
	//}

	go func() {
		for p := range ch {
			wg.Done()
			mp[p.i] = p.time
		}
	}()

	wg.Wait()

	i := 0
	for k, v := range mp {
		fmt.Println(k, v, i)
		i++
	}
}

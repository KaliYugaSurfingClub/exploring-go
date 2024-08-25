package fourth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
)

type R struct {
	Add int `json:"add"`
}

func getHandler(w http.ResponseWriter, counter *atomic.Int32) {
	w.Write([]byte(strconv.Itoa(int(counter.Load()))))
}

func postHandler(w http.ResponseWriter, r *http.Request, counter *atomic.Int32) {
	var data R

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < data.Add; i++ {
		counter.Add(1)
	}

	w.Write([]byte(strconv.Itoa(int(counter.Load()))))
}

func HttpMyExampleRace() {
	//no race condition
	counter := atomic.Int32{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getHandler(w, &counter)
		case http.MethodPost:
			postHandler(w, r, &counter)
		}
	})

	http.ListenAndServe(":8080", nil)
}

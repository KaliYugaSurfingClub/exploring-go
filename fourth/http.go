package fourth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UseHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//without writeHead also ok
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>GOOD<h1>"))
	})

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>page<h1>"))
	})

	http.ListenAndServe(":8080", nil)
}

func HttpMyExample() {
	counter := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(strconv.Itoa(counter)))
		case http.MethodPost:
			type R struct {
				Add int `json:"add"`
			}

			var data R

			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)

			if err != nil {
				fmt.Println(err)
			}

			//race condition
			//but no fatal errors with -race
			//cuz handler execute as a goroutine
			for i := 0; i < data.Add; i++ {
				counter++
			}

			w.Write([]byte(strconv.Itoa(counter)))
		}
	})

	http.ListenAndServe(":8080", nil)
}

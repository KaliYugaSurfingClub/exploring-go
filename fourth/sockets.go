package fourth

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// "Шина" событий, регистрация клиентов и рассылка сообщений идет отсюда
type Bus struct {
	register  chan *websocket.Conn
	broadcast chan []byte
	clients   map[*websocket.Conn]bool
}

func (b *Bus) Run() {
	for {
		select {
		case message := <-b.broadcast:
			// каждому зарегистрированному клиенту шлем сообщение
			for client := range b.clients {
				w, err := client.NextWriter(websocket.TextMessage)
				if err != nil {
					// если достучаться до клиента не удалось, то удаляем его
					delete(b.clients, client)
					continue
				}

				w.Write(message)
			}
		case client := <-b.register:
			// регистрируем клиентов в мапе клиентов
			log.Println("User registered")
			b.clients[client] = true
		}
	}
}

func NewBus() *Bus {
	return &Bus{
		register:  make(chan *websocket.Conn),
		broadcast: make(chan []byte),
		clients:   make(map[*websocket.Conn]bool),
	}
}

func runJoker(b *Bus) {
	for {
		// каждые 5 секунд ходим за шутками
		<-time.After(1 * time.Second)
		log.Println("Its joke time!")
		b.broadcast <- getJoke()
	}
}

func getJoke() []byte {
	return []byte(time.Now().String())
}

func UseSockets() {
	bus := NewBus()
	go bus.Run()
	go runJoker(bus)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//postman error
		//if r.Method == http.MethodGet {
		//	w.Write([]byte("123"))
		//}

		if r.Header.Get("Upgrade") == "websocket" {
			fmt.Println("New user")
			ws, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Fatal(err)
			}

			bus.register <- ws
		} else {
			w.Write([]byte("123"))
		}
	})

	http.ListenAndServe(":8080", nil)
}

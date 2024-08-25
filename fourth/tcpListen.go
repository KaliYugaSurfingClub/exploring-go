package fourth

import (
	"bufio"
	"fmt"
	"net"
)

func UseTcp() error {
	//use telnet 127.0.0.1 8080 to connect to this app
	listener, _ := net.Listen("tcp", ":8080")

	for {
		// ждём пока не придёт клиент
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Can not connect!!")
			conn.Close()
			continue
		}

		fmt.Println("Connected")

		// создаём Reader для чтения информации из сокета
		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		//defer conn.Close()

		go func(conn net.Conn) {
			for {
				//побайтово читаем
				rbyte, err := bufReader.ReadByte()

				if err != nil {
					fmt.Println("Can not read!", err)
					break
				}

				fmt.Print(string(rbyte))
				conn.Write([]byte{rbyte})
			}
		}(conn)
	}

}

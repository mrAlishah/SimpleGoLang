//24. Pool One Exercise Solution

//1- how to install netcat on ubuntu
//linux
//apt-get install nmap
//apt-get install nc

//MacOS
//brew install netcat
//brew install nmap
//brew install nc

//2- run go mainfile

//3- nmap 127.0.0.1 -p 1337,1338,1339

//4- nc 127.0.0.1 1337

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func handleConnections(port string, wg *sync.WaitGroup) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	fmt.Println("I'm lising port ", port)

LOOP:
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
		for {
			netData, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			temp := strings.TrimSpace(string(netData))
			if temp == "Stop" {
				break
			}

			conn.Write([]byte("Hi there form server... \n"))
		}

		conn.Close()
		break LOOP
	}

	fmt.Println("Closing Listener")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go handleConnections("1337", &wg)
	wg.Add(1)
	go handleConnections("1338", &wg)
	wg.Add(1)
	go handleConnections("1339", &wg)
	wg.Wait()

	fmt.Println("Server done handling connections")

}

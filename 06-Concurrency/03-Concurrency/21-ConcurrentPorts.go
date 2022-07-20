//26. Solution Concurrent Ports Concurrent Connections

//1- how to install netcat on ubuntu
//https://www.cyberithub.com/how-to-install-netcat-command-on-linux-ubuntu/
//https://zoomadmin.com/HowToInstall/UbuntuPackage/netcat
//sudo apt-get update -y
//sudo apt-get install -y netcat
//check out
//dpkg -L netcat
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

/*
 ➜  ~ nmap 127.0.0.1 -p 1337,1338,1339
Starting Nmap 7.92 ( https://nmap.org ) at 2022-04-14 16:53 +0430
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00014s latency).

PORT     STATE SERVICE
1337/tcp open  waste
1338/tcp open  wmc-log-svc
1339/tcp open  kjtsiteserver

Nmap done: 1 IP address (1 host up) scanned in 0.06 seconds
➜  ~ nc 127.0.0.1 1337
hi
Hi there form server...
dfd
Hi there form server...
kjkj
Hi there form server...
Stop
➜  ~ nmap 127.0.0.1 -p 1337,1338,1339
Starting Nmap 7.92 ( https://nmap.org ) at 2022-04-14 16:55 +0430
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00012s latency).

PORT     STATE  SERVICE
1337/tcp closed waste
1338/tcp open   wmc-log-svc
1339/tcp open   kjtsiteserver

Nmap done: 1 IP address (1 host up) scanned in 0.06 seconds

*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

func handleData(conn net.Conn, l net.Listener) {
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
	l.Close()

}

func handleConnections(port string, wg *sync.WaitGroup) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	fmt.Println("I'm lising port ", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Listener killed...\n", err, "\n")
			break
		}

		go handleData(conn, l)
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

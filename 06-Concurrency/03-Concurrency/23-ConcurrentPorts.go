//31. Slim Connection Pool Part I
//32. Using Our Slim Connection Pool

//https://github.com/sokil/go-connection-pool
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// ConnectionPool is a thread safe list of net.Conn instances
type ConnectionPool struct {
	mutex sync.RWMutex
	list  map[int]net.Conn
}

// NewConnectionPool is the factory method to create new connection pool
func NewConnectionPool() *ConnectionPool {
	pool := &ConnectionPool{
		list: make(map[int]net.Conn),
	}

	return pool
}

// Add collection to pool
func (pool *ConnectionPool) Add(connection net.Conn) int {
	pool.mutex.Lock()
	nextConnectionId := len(pool.list)
	pool.list[nextConnectionId] = connection
	pool.mutex.Unlock()
	return nextConnectionId
}

// Get connection by id
func (pool *ConnectionPool) Get(connectionId int) net.Conn {
	pool.mutex.RLock()
	connection := pool.list[connectionId]
	pool.mutex.RUnlock()
	return connection
}

// Remove connection from pool
func (pool *ConnectionPool) Remove(connectionId int) {
	pool.mutex.Lock()
	delete(pool.list, connectionId)
	pool.mutex.Unlock()
}

// Size of connections pool
func (pool *ConnectionPool) Size() int {
	return len(pool.list)
}

// Range iterates over pool
func (pool *ConnectionPool) Range(callback func(net.Conn, int)) {
	pool.mutex.RLock()
	for connectionId, connection := range pool.list {
		callback(connection, connectionId)
	}
	pool.mutex.RUnlock()
}

//Additions...
func handleData(nc net.Conn) {
	for {
		netData, err := bufio.NewReader(nc).ReadString('\n')
		if err != nil { fmt.Println(err); return; }

		clientMsg := strings.TrimSpace(string(netData))
		fmt.Println("Client wrote: ",clientMsg)
		//nc.Write([]byte(clientMsg+"\n"))

		if clientMsg == "Quit" { nc.Close(); break; }

	}

}

func main() {
	socket, err := net.Listen("tcp","127.0.0.1:1337")
	if err != nil { fmt.Println(err) }
	fmt.Println("I'm lising port  1337")

	connectionPool := NewConnectionPool()

	go func(pool *ConnectionPool){

		for {
			c, _ := socket.Accept()
			cid := pool.Add(c)
			fmt.Println("New client join, ID: ",cid)

			size := pool.Size()
			fmt.Println("Pool size: ",size)

			go handleData(pool.Get(cid))
			
			pool.Range( func(targetConnection net.Conn, targetConnectionId int) {

				writer := bufio.NewWriter(targetConnection)

				if (targetConnectionId != cid) {
					writer.WriteString("Got new connection \n")
				} else if (targetConnectionId == cid) {
					writer.WriteString("Wellcome to the sytem \n")
				}

				writer.Flush()

			})//pool.Range



		}//for

	}(connectionPool)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGTSTP, syscall.SIGQUIT)

	for {
		<-c
		fmt.Println("\n KILL PROGRAM")
		break;
	}
}
//37. Pool Project P2 Server

//How to run t
//1- go build 26-PoolServer.go
//2- GODEBUG=gctrace=1 ./26-PoolServer
package main

import (
	//Native Packages

	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"

	//3rd partiy
	//go get -u github.com/mgutz/logxi/v1
	log "github.com/mgutz/logxi/v1"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	rn = rand.New(s)

	// create package variable for Logger interface
	logger log.Logger
)

const ReqDataSize = 1 * 1024 //1KB
type ClientReq struct {
	ID      uint
	ReqType int               //One of ReqX defined above
	Data    [ReqDataSize]byte //request specific encode data
	Size    int               //how many byte in Data
}

type Server interface {
	Start() error
	Stop()
}

type TCPServer struct {
	numReqs uint64
	port    string
	log     log.Logger
	s       *http.Server
}

func newTCPServer(port string) Server {
	srv := &TCPServer{port: port}
	srv.log = log.New("server")

	//srv.log.SetLevel(log.LevelDebug)
	srv.log.SetLevel(log.LevelInfo)

	//TODO: Create and configure http.server
	s := &http.Server{}
	//configure http server
	s.WriteTimeout = 500 * time.Microsecond
	s.ReadTimeout = 1000 * time.Millisecond
	s.Addr = fmt.Sprintf(":%v", srv.port) //listen on all interfaces
	s.Handler = srv
	srv.s = s //store a refrence to the http.server
	srv.log.Info("Port: ", s.Addr)
	return srv
}

func (srv *TCPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//safely increment a number from multiple goroutines... better mutex locks etc ...
	atomic.AddUint64(&srv.numReqs, 1)

	srv.log.Debug("TCPServer - message from ", r.RemoteAddr)

		dec := json.NewDecoder(r.Body)
		defer r.Body.Close()

		msg := &ClientReq{}

		dec.Decode(msg)
		//INFO: pretent we do some work on with there
		time.Sleep(time.Duration(rn.Intn(5)) * time.Microsecond)

}

func (srv *TCPServer) Start() error {
	if nil == srv {
		return fmt.Errorf("Start() called on nil TCPServer object")
	}

	srv.log.Info("Starting HTTP Server")

	//Used with stop server thing
	var err error
	go func() {
		err = srv.s.ListenAndServe()
	}()
	time.Sleep(200 * time.Microsecond)
	return err

}

//Stop listening and close all clients connection
func (srv *TCPServer) Stop() {
	if nil == srv {
		return
	}

	srv.log.Info("Stopping HTTP Server")
	srv.s.Close()

	//for server close thing
	srv.log.Info("Message Proccessed: ", srv.numReqs)

}
func main() {
	srv := newTCPServer("7000")

	err := srv.Start()
	if err != nil {
		log.Error("Failed to start TCPServer", err)
		return
	}

	d := 20 * time.Second
	fmt.Printf("Sleep for %v\n", d)
	time.Sleep(d)
	srv.Stop()
}

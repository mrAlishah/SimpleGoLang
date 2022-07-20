//34. Pool Project p1
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	requestsPerClient = 100
	maxBatchSize = (requestsPerClient / 10) * 2 //20% of total request

	ReqAdd = iota
	ReqAve = iota
	ReqRandom = iota
	ReqSpellCheck = iota
	ReqSearch = iota
)

//ReqDataSize is the max  bytes per clientReq.Data byte array
const ReqDataSize = 1 * 1024 //1KB

type ClientReq struct{
	ID uint
	ReqType int					//One of ReqX defined above
	Data [ReqDataSize]byte 		//request specific encode data
	Size int 					//how many byte in Data
}

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	var req *ClientReq
	msgLeft  := requestsPerClient
	var reqID uint

	for 0 < msgLeft {
		batch := r.Intn(maxBatchSize)
		fmt.Printf("Batch: %d \n",batch)
		if batch > msgLeft {
			batch = msgLeft
		}
		msgLeft -= batch
		fmt.Printf("msgLeft: %d \n",msgLeft)

		for i := 0; i < batch; i++ {
			req = &ClientReq{}
			reqID++
			req.ID = reqID
			req.Size = r.Intn(ReqDataSize)
			for y:=0; y < req.Size; y++ {
				req.Data[y] = byte(y + 1)
			}

			fmt.Println(req)  //send to server
		}

		// pause a bit  between batches
		// time.Sleep(time.Duration(r.Intn(200)) * time.Millisecond)
	}
}
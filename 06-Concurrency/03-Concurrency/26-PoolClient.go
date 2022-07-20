//36. Pool Project P2 Client
package main

import (
	//Native Packages
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	//3rd partiy
	//go get -u github.com/mgutz/logxi/v1
	log "github.com/mgutz/logxi/v1"
)

// create package variable for Logger interface
var logger log.Logger

const (
	requestsPerClient = 1000
	maxBatchSize      = (requestsPerClient / 10) * 2 //20% of total request

	ReqAdd        = iota
	ReqAve        = iota
	ReqRandom     = iota
	ReqSpellCheck = iota
	ReqSearch     = iota
)

//ReqDataSize is the max  bytes per clientReq.Data byte array
const ReqDataSize = 1 * 1024 //1KB

type ClientReq struct {
	ID      uint
	ReqType int               //One of ReqX defined above
	Data    [ReqDataSize]byte //request specific encode data
	Size    int               //how many byte in Data
}

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func encodeReq(req *ClientReq) io.Reader {
	var buf = &bytes.Buffer{}
	jsonEnc := json.NewEncoder(buf)
	jsonEnc.Encode(req)
	return buf
}

func submitRequests(url string) {
	var req *ClientReq
	msgLeft := requestsPerClient
	var reqID uint

	for 0 < msgLeft {
		batch := r.Intn(maxBatchSize)
		fmt.Printf("Batch: %d \n", batch)
		if batch > msgLeft {
			batch = msgLeft
		}
		msgLeft -= batch
		fmt.Printf("msgLeft: %d \n", msgLeft)

		for i := 0; i < batch; i++ {
			req = &ClientReq{}
			reqID++
			req.ID = reqID
			req.Size = r.Intn(ReqDataSize)
			for y := 0; y < req.Size; y++ {
				req.Data[y] = byte(y + 1)
			}

			buf := encodeReq(req)
			resp, err := http.Post(url, "text/json", buf)
			if err != nil {
				log.Info("Hello: just test pkg")

				// create a logger with a unique identifier which
				// can be enabled from environment variables
				logger = log.New("pkg")
				logger.Error("Post error: ", err)
				break //try again later
			}
			defer resp.Body.Close()

			fmt.Println(req) //send to server
		}

		// pause a bit  between batches
		//time.Sleep(time.Duration(r.Intn(200)) * time.Millisecond)
	}
}
func main() {
	submitRequests("http://localhost:7000")
}

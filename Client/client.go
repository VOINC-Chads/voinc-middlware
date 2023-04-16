package main

import (
	"flag"
	"log"
	"strconv"

	"voinc/messages"

	"github.com/golang/protobuf/proto"
	"github.com/zeromq/goczmq"
)

func main() {
	var port int
	var addr string
	var requirements string
	var processCode string
	var executeCode string

	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.StringVar(&addr, "a", "localhost", "Provide an address for VOINC master")
	flag.StringVar(&requirements, "r", "req", "Provide a requirements.txt file")
	flag.StringVar(&processCode, "c", "process", "Provide a process code")
	flag.StringVar(&executeCode, "e", "execute", "Provide an execute code")

	flag.Parse()

	// Create a dealer socket and connect it to the router.
	dealer, err := goczmq.NewDealer("tcp://" + addr + ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()

	log.Println("dealer created and connected")

	mainMsg := &messages.MainReq{
		MsgType: messages.MsgTypes_TYPE_CODE,
		Content: &messages.MainReq_CodeMsg{
			CodeMsg: &messages.CodeMsg{
				Requirements: requirements,
				ProcessCode:  processCode,
				ExecuteCode:  executeCode,
			},
		},
	}

	protoMsg, err := proto.Marshal(mainMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Send a 'Hello' message from the dealer to the router.
	// Here we send it as a frame ([]byte), with a FlagNone
	// flag to indicate there are no more frames following.
	err = dealer.SendFrame([]byte(protoMsg), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("dealer sent '" + string(protoMsg) + "'")

	// Register a poller
	poller, err := goczmq.NewPoller()
	if err != nil {
		log.Fatal(err)
	}
	poller.Add(dealer)

	// Wait for a message
	for {
		socket := poller.Wait(-1)

		if socket == dealer {
			// Receive the message. Here we call RecvMessage, which
			// will return the message as a slice of frames ([][]byte).
			// Since this is a router socket that support async
			// request / reply, the first frame of the message will
			// be the routing frame.
			request, err := dealer.RecvMessage()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("dealer received '%s' from '%v'", request[1], request[0])
			break
		} else {
			log.Fatal("Unexpected socket")
		}
	}
}

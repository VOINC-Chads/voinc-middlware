package main

import (
	"flag"
	"io/ioutil"
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

	// Read from executeCode.py into executeCode variable
	fileContent, err := ioutil.ReadFile("./themscripts/executeCode.py")
	if err != nil {
		log.Fatal(err)
	}

	executeCode = string(fileContent)

	// Read from processCode.py into processCode variable
	fileContent, err = ioutil.ReadFile("./themscripts/processCode.py")
	if err != nil {
		log.Fatal(err)
	}

	processCode = string(fileContent)

	// Read from requirements.txt into requirements variable
	fileContent, err = ioutil.ReadFile("./themscripts/requirements.txt")
	if err != nil {
		log.Fatal(err)
	}

	requirements = string(fileContent)

	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.StringVar(&addr, "a", "localhost", "Provide an address for VOINC master")

	flag.Parse()

	// Create a dealer socket and connect it to the router.
	dealer, err := goczmq.NewDealer("tcp://" + addr + ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()

	log.Println("dealer created and connected")

	heartbeatMsg := &messages.MainReq{
		MsgType: messages.MsgTypes_TYPE_HEARTBEAT,
		Content: &messages.MainReq_Heartbeat{
			Heartbeat: &messages.Heartbeat{},
		},
	}

	protoMsg, err := proto.Marshal(heartbeatMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	err = dealer.SendFrame([]byte(protoMsg), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

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

			log.Printf("dealer received")
			log.Println(request)
			break
		} else {
			log.Fatal("Unexpected socket")
		}
	}

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

	protoMsg, err = proto.Marshal(mainMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	err = dealer.SendFrame([]byte(protoMsg), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("dealer sent '" + string(protoMsg) + "'")

	mainMsg = &messages.MainReq{
		MsgType: messages.MsgTypes_TYPE_JOB,
		Content: &messages.MainReq_JobMsg{
			JobMsg: &messages.JobMsg{
				Jobs: []string{"1", "2"},
			},
		},
	}

	protoMsg, err = proto.Marshal(mainMsg)
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

			log.Printf("dealer received")
			log.Println(request)
		} else {
			log.Fatal("Unexpected socket")
		}
	}
}

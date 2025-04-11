package client

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net"
	"socket-tcp-udp/pkg"
	"time"
)

func UdpRpcClient() {
	conn, err := net.DialTimeout("udp", "127.0.0.1:8888", 2*time.Minute)
	if err != nil {
		log.Fatalf("Failed to connect to UDP server: %v", err)
	}
	defer conn.Close()

	request := pkg.Request{
		RequestId: rand.Int(),
		Name:      "world",
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Failed to marshal request: %v", err)
	}
	_, err = conn.Write(requestData)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}

	response := make([]byte, 256)
	n, err := conn.Read(response)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	var resp pkg.Response
	err = json.Unmarshal(response[:n], &resp)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	log.Printf("Received response: %s", resp.SayHello)

}

package server

import (
	"encoding/json"
	"socket-tcp-udp/pkg"

	"log"
	"net"
)

func UdpRpcServer() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Failed to start UDP listener: %v", err)
	}
	defer conn.Close()

	request := make([]byte, 256)
	n, remoteAddr, err := conn.ReadFromUDP(request)
	if err != nil {
		log.Fatalf("Error reading from connection: %v", err)
	}
	response := handle(request[:n])
	if len(response) > 0 {
		_, err = conn.WriteToUDP(response, remoteAddr)
		if err != nil {
			log.Fatalf("Error sending response: %v", err)
		}
		log.Printf("Sent response to %s: %s", remoteAddr, string(response))
	}
}

func handle(request []byte) (response []byte) {
	var req pkg.Request
	if err := json.Unmarshal(request, &req); err != nil {
		log.Printf("Failed to unmarshal request: %v", err)
		return nil
	}

	resp := pkg.Response{RequestId: req.RequestId, SayHello: "Hello " + req.Name}
	response, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
		return nil
	}
	return response
}

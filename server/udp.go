package server

import (
	"log"
	"net"
	"sync"
)

func UdpServer() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5678")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Failed to start UDP listener: %v", err)
	}
	defer conn.Close()
	log.Println("Waiting for connection...")

	request := make([]byte, 256)
	n, remoteAddr, err := conn.ReadFromUDP(request)
	if err != nil {
		log.Fatalf("Error reading from connection: %v", err)
	}
	log.Printf("Received request from %s: %s", remoteAddr, string(request[:n]))

	response := []byte("world")
	_, err = conn.WriteToUDP(response, remoteAddr)
	if err != nil {
		log.Fatalf("Error sending response: %v", err)
	}
	log.Printf("Sent response to %s: %s", remoteAddr, string(response))
}

func UdpConnectionCurrent() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:3333")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Failed to start UDP listener: %v", err)
	}
	defer conn.Close()
	log.Println("Waiting for connection...")

	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			request := make([]byte, 256)
			n, remoteAddr, err := conn.ReadFromUDP(request)
			if err != nil {
				log.Fatalf("Error reading from connection: %v", err)
			}
			log.Printf("Received request from %s: %s", remoteAddr, string(request[:n]))
		}()
	}
	wg.Wait()

}

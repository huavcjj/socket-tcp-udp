package client

import (
	"log"
	"net"
	"sync"
	"time"
)

func connectToUdpServer(serverAddr string) net.Conn {
	conn, err := net.DialTimeout("udp", serverAddr, 2*time.Minute)
	if err != nil {
		log.Fatalf("Failed to connect to UDP server: %v", err)
	}
	log.Printf("Connected to UDP server: remote=%s, local=%s", conn.RemoteAddr(), conn.LocalAddr())
	return conn
}

func sendRequestToUdpServer(conn net.Conn) {
	n, err := conn.Write([]byte("hello"))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	log.Printf("Sent request: %d bytes", n)
}

func UdpClient() {
	conn := connectToUdpServer("127.0.0.1:5678")
	defer conn.Close()

	sendRequestToUdpServer(conn)

	response := make([]byte, 256)
	n, err := conn.Read(response)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	log.Printf("Received response: %s", string(response[:n]))

}

func UdpConnectionCurrent() {
	conn := connectToUdpServer("127.0.0.1:3333")
	defer conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			sendRequestToUdpServer(conn)
		}()

	}
	wg.Wait()
}

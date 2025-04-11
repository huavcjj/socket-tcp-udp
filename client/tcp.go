package client

import (
	"encoding/binary"
	"log"
	"net"
	"time"
)

func connectToTcpServer(serverAddr string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to resolve TCP address: %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	return conn
}

func sendRequestToTcpServer(conn net.Conn) error {
	n, err := conn.Write([]byte("hello"))
	if err != nil {
		return err
	}
	log.Println("Sent request:", n)
	return nil
}

func TcpClient() {
	conn := connectToTcpServer("127.0.0.1:5678")
	defer conn.Close()

	if err := sendRequestToTcpServer(conn); err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
}

func SendMultipleTCPRequests() {
	conn := connectToTcpServer("127.0.0.1:8080")
	defer conn.Close()

	log.Println("Connected to server")

	for i := 0; i < 10; i++ {
		if err := sendRequestToTcpServer(conn); err != nil {
			log.Fatalf("Request %d failed: %v", i+1, err)
		}
		log.Printf("Request %d sent", i+1)
		time.Sleep(100 * time.Millisecond)
	}

	log.Println("All requests sent")
}

func sendFramedMessage(conn net.Conn, message string) error {
	payload := []byte(message)
	length := uint32(len(payload))

	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, length)

	_, err := conn.Write(append(header, payload...))
	return err
}

func RunFramedClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	messages := []string{"hello", "world", "this is framed", "by length header"}
	for i, msg := range messages {
		if err := sendFramedMessage(conn, msg); err != nil {
			log.Fatalf("Failed to send message #%d: %v", i+1, err)
		}
		log.Printf("Sent message #%d: %s", i+1, msg)
	}
}

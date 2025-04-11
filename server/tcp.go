package server

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
)

func TcpServer() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5678")
	if err != nil {
		log.Fatalf("Failed to resolve TCP address: %v", err)
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Fatalf("Failed to start TCP listener: %v", err)
	}
	defer listener.Close()
	log.Println("Waiting for connection...")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	defer conn.Close()
	log.Printf("Established connection from %s to %s", conn.RemoteAddr(), conn.LocalAddr())

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))

	request := make([]byte, 256)
	n, err := conn.Read(request)
	if err != nil {
		log.Printf("Error reading from connection: %v", err)
		return
	}
	log.Printf("Received request: %s", string(request[:n]))
}

func HandleLongTCPConnection() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to resolve TCP address: %v", err)
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Fatalf("Failed to start TCP listener: %v", err)
	}
	defer listener.Close()
	log.Println("Waiting for connection...")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	defer conn.Close()
	log.Printf("Established connection from %s to %s", conn.RemoteAddr(), conn.LocalAddr())

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	request := make([]byte, 256)
	for {
		n, err := conn.Read(request)
		if err == io.EOF {
			log.Println("Client closed the connection (EOF received)")
			break
		}
		if err != nil {
			log.Printf("Error reading from connection: %v", err)
			break
		}
		log.Printf("Received request: %s", string(request[:n]))

		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	}
}

func TcpStick() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:5555")
	if err != nil {
		log.Fatalf("Failed to resolve TCP address: %v", err)
	}

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		log.Fatalf("Failed to start TCP listener: %v", err)
	}
	log.Println("Waiting for client connection...")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	defer conn.Close()
	log.Printf("Established connection from %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		header := make([]byte, 4)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			if err == io.EOF {
				log.Println("Client closed the connection")
			} else {
				log.Printf("Error reading length header: %v", err)
			}
			break
		}

		length := binary.BigEndian.Uint32(header)
		if length == 0 {
			continue
		}

		payload := make([]byte, length)
		_, err = io.ReadFull(reader, payload)
		if err != nil {
			log.Printf("Error reading payload: %v", err)
			break
		}

		log.Printf("Received message: %s", string(payload))
	}
}

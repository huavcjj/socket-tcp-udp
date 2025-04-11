package server_test

import (
	"socket-tcp-udp/server"
	"testing"
)

func TestTcpServer(t *testing.T) {
	server.TcpServer()
}

func TestUdpServer(t *testing.T) {
	server.UdpServer()
}

func TestHandleLongTCPConnection(t *testing.T) {
	server.HandleLongTCPConnection()
}

func TestTcpStick(t *testing.T) {
	server.TcpStick()
}

func TestUdpConnectionCurrent(t *testing.T) {
	server.UdpConnectionCurrent()
}

func TestUdpRpcServer(t *testing.T) {
	server.UdpRpcServer()
}

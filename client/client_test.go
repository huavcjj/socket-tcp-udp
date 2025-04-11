package client_test

import (
	"socket-tcp-udp/client"
	"testing"
)

func TestTcpClient(t *testing.T) {
	client.TcpClient()
}

func TestUdpClient(t *testing.T) {
	client.UdpClient()
}

func TestSendMultipleTCPRequests(t *testing.T) {
	client.SendMultipleTCPRequests()
}

func TestRunFramedClient(t *testing.T) {
	client.RunFramedClient()
}

func TestUdpConnectionCurrent(t *testing.T) {
	client.UdpConnectionCurrent()
}

func TestUdpRpcClient(t *testing.T) {
	client.UdpRpcClient()
}

package handler

import (
	"log"
	"net"
	"os"
)

const (
	CChanLen  = 1000
	CDataSize = 1024
)

var LimitChan = make(chan string, CChanLen)
var doneChan = make(chan bool, CChanLen)

func UdpServer(port int) {
	log.Println("start listening on port:", port)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	})

	defer conn.Close()
	if err != nil {
		log.Println("read from connect failed, err:" + err.Error())
		os.Exit(1)
	}

	for {
		doneChan <- true
		udpProcess(conn)
	}
}

func udpProcess(conn *net.UDPConn) {
	// 最大读取数据大小
	data := make([]byte, CDataSize)
	n, address, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Println("failed read udp msg, error: " + err.Error())
	}
	log.Println("received adddress:", *address)
	str := string(data[:n])
	log.Println("receive from client, data:" + str)
	LimitChan <- str
	<-doneChan
}

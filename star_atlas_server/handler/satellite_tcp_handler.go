package handler

import (
	"fmt"
	"net"
	"star_atlas_server/pb"

	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

const cBufferSize = 102400

func SatelliteTCPHandler(tcpPort int) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpPort))
	if err != nil {
		glog.Errorf("Failed to listen tcp port: %v err: %v\n", tcpPort, err)
		return
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		glog.Errorf("Failed to accept connection: %v err: %v\n", l.Addr().String(), err)
	}
	defer conn.Close()

	for {
		data, err := readBytes(conn)
		if err != nil {
			glog.Errorf("Failed to read from connection: %v err: %v\n", conn, err)
			continue
		}
		handler(conn, data)
	}
}

func handler(conn net.Conn, data []byte) error {
	data = data[4:]
	m := &pb.Msg{}
	if err := proto.Unmarshal(data, m); err != nil {
		glog.Errorf("Failed to unmarshal message: %v\n", err)
		return err
	}
	//todo handler msg then send result
	n, err := conn.Write(nil)
	if err != nil {
		glog.Errorf("Failed to write connection message: %v\n", err)
		return err
	}
	glog.Infof("Connection message written by length %d bytesData\n", n)
	return nil
}

func readBytes(conn net.Conn) ([]byte, error) {
	tmp := make([]byte, cBufferSize)
	n, err := conn.Read(tmp)
	if err != nil {
		return nil, err
	}
	return tmp[:n], nil
}

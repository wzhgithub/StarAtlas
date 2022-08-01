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

// ==============================todo case implementations =============================
func handleByApiType(msg *pb.Msg) ([]byte, error) {
	switch msg.Type {
	case pb.MsgType_ApiOrbitNormal:
	case pb.MsgType_ApiShowPicture:
	case pb.MsgType_ApiMarkerCoordinates:
	case pb.MsgType_ApiOrbitCoordinate:
	}

	return make([]byte, 0), nil
}

func handler(conn net.Conn, data []byte) error {
	data = data[4:]
	m := &pb.Msg{}
	if err := proto.Unmarshal(data, m); err != nil {
		glog.Errorf("Failed to unmarshal message: %v\n", err)
		return err
	}
	//todo handler msg then send result
	res, err := handleByApiType(m)
	if err != nil {
		glog.Errorf("Failed to handler message: %v\n", err)
		return err
	}
	n, err := conn.Write(res)
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

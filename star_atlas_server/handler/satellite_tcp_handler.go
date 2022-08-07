package handler

import (
	"encoding/binary"
	"fmt"
	"net"
	"star_atlas_server/model"
	"star_atlas_server/pb"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"
)

var conn net.Conn

type ToProto interface {
	ToProto() (proto.Message, error)
}

type OrbitNormal struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type Coordinate struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

type OrbitCoordinate struct {
	Coordinates []*Coordinate `json:"coordinates"`
}

func (O *OrbitCoordinate) GetType() (pb.MsgType, error) {
	if O.Coordinates == nil || len(O.Coordinates) == 0 {
		return 0, fmt.Errorf("invalidate coordinates")
	}
	lc := len(O.Coordinates)
	if lc == 1 {
		return pb.MsgType_ApiOrbitCoordinate, nil
	}

	return pb.MsgType_ApiMarkerCoordinates, nil
}

func (O *OrbitCoordinate) ToProto() (proto.Message, error) {
	if O.Coordinates == nil || len(O.Coordinates) == 0 {
		return nil, fmt.Errorf("invalidate coordinates")
	}
	lc := len(O.Coordinates)
	if lc == 1 {
		return &pb.OrbitCoordinate{
			Coord: &pb.Coordinate{
				Longitude: O.Coordinates[0].Longitude,
				Latitude:  O.Coordinates[1].Latitude,
			},
		}, nil
	}
	coordinates := &pb.MarkerCoordinates{
		Coordinates: make([]*pb.Coordinate, 0),
	}
	for i := 0; i < lc; i++ {
		coordinates.Coordinates = append(coordinates.Coordinates, &pb.Coordinate{
			Longitude: O.Coordinates[i].Longitude,
			Latitude:  O.Coordinates[i].Latitude,
		})
	}
	return coordinates, nil
}

func (O *OrbitNormal) ToProto() (proto.Message, error) {
	return &pb.OrbitNormal{
		X: O.X,
		Y: O.Y,
		Z: O.Z,
	}, nil
}

func ApiOrbitNormal(c *gin.Context) {
	o := &OrbitNormal{}
	if err := c.ShouldBindJSON(o); err != nil {
		glog.Errorf("ApiOrbitNormal Error binding JSON: %v", err)
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	msg, err := o.ToProto()
	if err != nil {
		glog.Errorf("ApiOrbitNormal Error ToProto: %v", err)
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	err = sendMsg(conn, pb.MsgType_ApiOrbitNormal, msg)
	if err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	c.JSON(200, model.NewCommonResponseSucc("ApiOrbitNormal success"))
}

func apiCoordinate(c *gin.Context) {
	req := &OrbitCoordinate{}
	if err := c.ShouldBindJSON(req); err != nil {
		glog.Errorf("apiCoordinate Error binding JSON: %v", err)
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	msg, err := req.ToProto()
	if err != nil {
		glog.Errorf("apiCoordinate Error ToProto: %v", err)
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	t, _ := req.GetType()
	err = sendMsg(conn, t, msg)
	if err != nil {
		c.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	c.JSON(200, model.NewCommonResponseSucc("ApiOrbitCoordinate success"))
}

func ApiMarkerCoordinates(c *gin.Context) {
	apiCoordinate(c)
}

func ApiOrbitCoordinate(c *gin.Context) {
	apiCoordinate(c)
}

func ApiShowPicture(ctx *gin.Context) {
	name := ctx.Query("name")
	msg := &pb.ShowPicture{
		Name: name,
	}
	err := sendMsg(conn, pb.MsgType_ApiShowPicture, msg)
	if err != nil {
		ctx.JSON(500, model.NewCommonResponseFail(err))
		return
	}
	ctx.JSON(200, model.NewCommonResponseSucc("ApiShowPicture success"))
}

func SatelliteTCPHandlerInit(tcpPort int) {
	glog.Infof("SatelliteTCPHandlerInit called port:%d\n", tcpPort)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpPort))
	if err != nil {
		glog.Errorf("Failed to listen tcp port: %v err: %v\n", tcpPort, err)
		return
	}
	conn, err = l.Accept()
	if err != nil {
		glog.Errorf("Failed to accept connection: %v err: %v\n", l.Addr().String(), err)
	}
	glog.Infof("Connected to %s", conn.LocalAddr().String())
	for {
		data := make([]byte, 4096*1000)
		n, err := conn.Read(data)
		if err != nil {
			glog.Errorf("Error reading %s: %v\n", conn.LocalAddr().String(), err)
			continue
		}
		data = data[:n]
		go handleMsg(conn, data)
	}
}

func handleMsg(conn net.Conn, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("go handle tcp Msg error: %v\n", err)
		}
	}()
	msg := &pb.Msg{}
	err := proto.Unmarshal(data, msg)
	if err != nil {
		glog.Errorf("go Unmarshal tcp Msg error: %v\n", err)
		return
	}
	glog.Infof("go Unmarshal tcp Msg: %v\n", msg)
}

func sendMsg(conn net.Conn, msgType pb.MsgType, msg proto.Message) error {
	if conn == nil {
		return fmt.Errorf("connection must be non-nil")
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	data, err = proto.Marshal(&pb.Msg{
		Type: msgType,
		Data: data,
	})
	if err != nil {
		return err
	}

	var writeBuf = make([]byte, 4)
	binary.BigEndian.PutUint32(writeBuf, uint32(len(data)))
	writeBuf = append(writeBuf, data...)
	n, err := conn.Write(writeBuf)
	if n != len(writeBuf) || err != nil {
		return err
	}
	return nil
}

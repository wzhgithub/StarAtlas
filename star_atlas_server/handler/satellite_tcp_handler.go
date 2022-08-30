package handler

import (
	"encoding/binary"
	"fmt"
	"math"
	// "math/rand"
	"net"
	"star_atlas_server/config"
	"star_atlas_server/model"
	"star_atlas_server/pb"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/nl8590687/asrt-sdk-go/sdk"
	"google.golang.org/protobuf/proto"
)

var conn net.Conn

var picCnt = 0

var posMap = map[string]string{"san": "W", "shan": "W", "shang": "W", "sang": "W", "xia": "S", "zhuo": "A", "zuo": "A", "yong": "D", "you": "D"}

const cBufferSize = 4096 * 1000
const cStep = math.Pi / 36 // 5 dgree

var basePolar = &Polar{
	Theta: 0,
	Phi:   0,
	R:     1,
}

type Polar struct {
	Theta float64
	Phi   float64
	R     float64
}

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

func (o *Polar) OperationByKey(keyName string) (*OrbitNormal, error) {
	/**
	x = sin(theta)*cos(phi) 0
	y = sin(theta)*sin(phi) 0
	z = cos(theta)  1
	theta = 0
	phi = 0
	*/

	switch keyName {
	case "W":
		o.Theta += cStep
	case "S":
		o.Theta -= cStep
	case "A":
		o.Phi += cStep
	case "D":
		o.Phi -= cStep
	default:
		return nil, fmt.Errorf("keyName not supported: %s", keyName)
	}
	glog.Infof("basePolar: %v\n", o)

	orb := &OrbitNormal{
		X: float32(o.R) * float32(math.Sin(o.Theta)) * float32(math.Cos(o.Phi)),
		Z: float32(o.R) * float32(math.Sin(o.Theta)) * float32(math.Sin(o.Phi)),
		Y: float32(o.R) * float32(math.Cos(o.Theta)),
	}

	return orb, nil
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

func newConn(tcpPort int) (net.Conn, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpPort))
	if err != nil {
		glog.Errorf("Failed to listen tcp port: %v err: %v\n", tcpPort, err)
		return nil, err
	}
	return l.Accept()
}

func SatelliteTCPHandlerInit(tcpPort int) {
	glog.Infof("SatelliteTCPHandlerInit called port:%d\n", tcpPort)
	var err error
	conn, err = newConn(tcpPort)
	if err != nil {
		glog.Errorf("failed connection e:%v\n", err)
		return
	}
	glog.Infof("Connected to %s", conn.LocalAddr().String())
	for {
		data := make([]byte, cBufferSize)
		n, err := conn.Read(data)
		if err != nil {
			glog.Errorf("Error reading %s: %v\n", conn.LocalAddr().String(), err)
			conn.Close()
			conn, _ = newConn(tcpPort)
			continue
		}
		data = data[:n]
		handleMsg(conn, data)
	}
}

func handleMsg(conn net.Conn, data []byte) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("go handle tcp Msg error: %v\n", err)
		}
	}()
	l := binary.BigEndian.Uint32(data[0:4])
	glog.Infof("received start:%d end:%d\n", 4, 4+l)
	if l > cBufferSize {
		glog.Warningf("received out of range data len:%d buffer:%d\n", l, cBufferSize)
		return
	}
	msgData := data[4 : 4+l]
	msg := &pb.Msg{}
	err := proto.Unmarshal(msgData, msg)
	if err != nil {
		glog.Errorf("go Unmarshal tcp Msg error: %v\n", err)
		return
	}
	glog.Infof("go Unmarshal tcp Msg: %v\n", msg.GetType())
	if err = handlePbMsg(msg); err != nil {
		glog.Errorf("go handlePbMsg failed err: %v\n", err)
		return
	}
}

func handlePbMsg(msg *pb.Msg) error {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("go handlePbMsg error: %v", err)
		}
	}()
	switch msg.GetType() {
	case pb.MsgType_ApiExit:
		if conn != nil {
			glog.Infof("conn close: %v\n", conn)
			conn.Close()
		}
		picCnt = 0
		var err error
		conn, err = newConn(config.CommonConfig.SatelliteTCPPort)
		if err != nil {
			return err
		}
	case pb.MsgType_ApiSpeech:
		wave, err := sdk.DecodeWav(msg.Data)
		if err != nil {
			return err
		}
		// glog.Infof("wave bytes:%v\n", wave.GetRawSamples())
		res, err := RecogniteByType(wave.GetRawSamples(), 16000, 1, 2, CSpeechType)
		if err != nil {
			return err
		}
		glog.Infof("received speech from server %v\n", res)
		if res.StatusCode != 200000 {
			return fmt.Errorf(res.StatucMesaage)
		}
		glog.Infof("received speech word %v\n", res.Result)
		p := parseWAVCmd(res.Result.([]interface{}))
		if p == "" {
			return fmt.Errorf("received speech word not found")
		}
		return handleKeyboardOri(p)

	case pb.MsgType_ApiKeyboard:
		return handleKeyboardMessage(msg)
	}

	return nil
}

func handleKeyboardOri(key string) error {
	ori, err := basePolar.OperationByKey(key)
	if err != nil {
		return err
	}
	oMsg, err := ori.ToProto()
	if err != nil {
		return err
	}
	glog.Infof("key:%v OrbitNormal:%v\n", key, oMsg)
	if err = sendMsg(conn, pb.MsgType_ApiOrbitNormal, oMsg); err != nil {
		return err
	}
	return nil
}

func handleKeyboardMessage(msg *pb.Msg) error {
	k := &pb.Key{}
	if err := proto.Unmarshal(msg.GetData(), k); err != nil {
		return err
	}
	if k.GetName() == "Return" {
		return handlePic()
	}
	return handleKeyboardOri(k.GetName())
}

func handlePic() error {
	idx := picCnt % 10
	imageName := fmt.Sprintf("港口%d", idx)
	msg := &pb.ShowPicture{
		Name: imageName,
	}
	glog.Infof("pic data msg:%v\n", msg)
	err := sendMsg(conn, pb.MsgType_ApiShowPicture, msg)
	if err != nil {
		return err
	}
	picCnt += 1
	return nil
}

func parseWAVCmd(result []interface{}) string {
	for _, v := range result {
		for k, val := range posMap {
			vs := fmt.Sprint(v)
			glog.Infof("debug vs:%s and key:%s value:%s\n", vs, k, val)
			if strings.Contains(vs, val) {
				return k
			}
		}
	}
	return ""
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

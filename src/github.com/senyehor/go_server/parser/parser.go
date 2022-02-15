package parser

import (
	"github.com/senyehor/go_server/utils"
	"strconv"
	"strings"
)

const (
	PacketValuesCount = 16
)

var (
	packetConfig        = utils.GetPacketConfig()
	packetStructIndexes = PacketStructIndexes{
		tokenIndex:     0,
		valuesIndexes:  [2]uint8{1, PacketValuesCount + 1},
		timeIndex:      1 + PacketValuesCount,
		packetNumIndex: 2 + PacketValuesCount,
		deviceIDIndex:  3 + PacketValuesCount,
	}
)

type Packet struct {
	values    [PacketValuesCount]float64
	time      int
	packetNum int
	deviceID  int
}

func (p Packet) Values() *[PacketValuesCount]float64 {
	return &p.values
}

func (p Packet) Time() int {
	return p.time
}

func (p Packet) PacketNum() int {
	return p.packetNum
}

func (p Packet) DeviceID() int {
	return p.deviceID
}

type PacketStructIndexes struct {
	tokenIndex     uint8
	valuesIndexes  [2]uint8
	timeIndex      uint8
	packetNumIndex uint8
	deviceIDIndex  uint8
}

func NewPacket(binaryData []byte) *Packet {
	// returns nil if parsing binaryData goes wrong otherwise packet obj
	// [Token] ; [n1] ; [n2] ; ... ; [npacketConfig.ValuesCount()] ; [Time] ; [NUMpacket] ; [IDdevice]! - packet structure
	if (len(binaryData)) == 0 {
		return nil
	}
	packetParts := strings.Split(string(binaryData[:]), packetConfig.DataDelimiter())
	// 4 is non-data packetParts count
	if len(packetParts) != 4+PacketValuesCount {
		return nil
	}
	// deleting data terminator
	packetParts[packetStructIndexes.deviceIDIndex] = strings.
		TrimRight(packetParts[packetStructIndexes.deviceIDIndex], string(packetConfig.DataTerminator()))
	if strings.Compare(packetParts[packetStructIndexes.tokenIndex], packetConfig.Token()) != 0 {
		return nil
	}
	var values [PacketValuesCount]float64
	for i := packetStructIndexes.valuesIndexes[0]; i < packetStructIndexes.valuesIndexes[1]; i++ {
		parsedValue, err := strconv.ParseFloat(packetParts[i], 32)
		if err != nil {
			return nil
		}
		values[i-1] = parsedValue
	}
	time, err := strconv.Atoi(packetParts[packetStructIndexes.timeIndex])
	if err != nil {
		return nil
	}
	packetNum, err := strconv.Atoi(packetParts[packetStructIndexes.packetNumIndex])
	if err != nil {
		return nil
	}
	deviceID, err := strconv.Atoi(packetParts[packetStructIndexes.deviceIDIndex])
	if err != nil {
		return nil
	}
	return &Packet{
		values:    values,
		time:      time,
		packetNum: packetNum,
		deviceID:  deviceID,
	}

}

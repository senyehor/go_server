package binary_parser

import (
	"errors"
	"github.com/senyehor/go_server/data_models"
)

// ParseFromBinary returns nil if parsing goes wrong otherwise *packet
func ParseFromBinary(binaryData []byte) (*data_models.Packet, error) {
	packetParts, err := parseBinaryDataToStringParts(binaryData)
	if err != nil {
		return nil, errors.New("failed to split packet data to parts")
	}

	values, err := parsePacketValues(packetParts.Values())
	if err != nil {
		return nil, errors.New("failed to parse value from sensor")
	}
	time, err := parsePacketTimeInterval(packetParts.Time())
	if err != nil {
		return nil, errors.New("failed to parse timeInterval")
	}
	number, err := parsePacketNumber(packetParts.PacketNumber())
	if err != nil {
		return nil, errors.New("failed to parse packet number")
	}
	deviceID, err := parsePacketDeviceID(packetParts.DeviceID())
	if err != nil {
		return nil, errors.New("failed to parse device id")
	}

	packet, err := data_models.NewPacket(values, time, number, deviceID)
	if err != nil {
		return nil, errors.New("failed to create packet")
	}
	return packet, nil
}

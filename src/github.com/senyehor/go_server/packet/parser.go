package packet

import (
	"errors"
)

func ParseFromBinary(binaryData []byte) (*Packet, error) {
	// function returns nil if parsing goes wrong otherwise Packet obj
	packetParts, err := parseBinaryDataToStringParts(binaryData)
	if err != nil {
		return nil, err
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
		return nil, errors.New("failed to parse Packet number")
	}
	deviceID, err := parsePacketDeviceID(packetParts.DeviceID())
	if err != nil {
		return nil, errors.New("failed to parse device id")
	}

	return NewPacket(values, time, number, deviceID), nil
}

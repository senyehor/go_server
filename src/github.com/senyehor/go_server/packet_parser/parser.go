package packet_parser

import (
	"errors"
)

var (
	parsedDataPacketIndexes = getPacketPartsIndexesInParsedData()
)

func ParseFromBinary(binaryData []byte) (*Packet, error) {
	// returns nil if parsing goes wrong otherwise packet obj
	packetParts, err := parseBinaryDataToStringParts(binaryData)
	if err != nil {
		return nil, err
	}

	values, err := parsePacketValues(packetParts)
	if err != nil {
		return nil, errors.New("failed to parse value from sensor")
	}
	time, err := parsePacketTime(packetParts)
	if err != nil {
		return nil, errors.New("failed to parse time")
	}
	number, err := parsePacketNumber(packetParts)
	if err != nil {
		return nil, errors.New("failed to parse packet number")
	}
	deviceID, err := parsePacketDeviceID(packetParts)
	if err != nil {
		return nil, errors.New("failed to parse device id")
	}

	return &Packet{
		values:    values,
		time:      time,
		packetNum: number,
		deviceID:  deviceID,
	}, nil
}

package packet_parser

import (
	"errors"
	"github.com/senyehor/go_server/utils"
	"strconv"
	"strings"
)

// todo think of renaming file

var (
	packetConfig = utils.GetPacketConfig()
)

func getPacketPartsIndexesInParsedData() *packetPartsIndexesInParsedData {
	// [Token];[n1];[n2];...;[packetConfig.ValuesCount()];[Time];[PacketNumber];[IDdevice]!
	//- packet structure
	return &packetPartsIndexesInParsedData{
		token:              0,
		valuesRangeBorders: rangeBorders{1, PacketValuesCount + 1}, // left border included, second excluded
		// indexes below are dependent on PacketValuesCount
		//and each shifts to one more from right border of values right border
		time:         1 + PacketValuesCount,
		packetNumber: 2 + PacketValuesCount,
		deviceID:     3 + PacketValuesCount,
	}
}

func parseBinaryDataToStringParts(binaryData []byte) ([]string, error) {
	parts := strings.Split(string(binaryData[:]), string(packetConfig.DataDelimiter()))
	if !checkPacketLength(parts) {
		return parts, errors.New("invalid packet length")
	}
	if !checkPacketToken(parts) {
		return parts, errors.New("invalid packet token")
	}
	// deleting data terminator
	parts[parsedDataPacketIndexes.deviceID] =
		strings.TrimRight(
			parts[parsedDataPacketIndexes.deviceID],
			string(packetConfig.DataTerminator()),
		)
	return parts, nil
}

func checkPacketLength(packetParts []string) bool {
	return len(packetParts) == NonValuesPacketPartsCount+PacketValuesCount
}

func checkPacketToken(packetParts []string) bool {
	return strings.Compare(packetParts[parsedDataPacketIndexes.token], packetConfig.Token()) == 0
}

func parsePacketValues(packetParts []string) (packetValues, error) {
	var values packetValues
	for partsIndexCounter, valuesIndexCounter := parsedDataPacketIndexes.valuesRangeBorders.left, 0; partsIndexCounter < parsedDataPacketIndexes.valuesRangeBorders.right; partsIndexCounter, valuesIndexCounter = partsIndexCounter+1, valuesIndexCounter+1 {
		parsedValue, err := strconv.ParseFloat(packetParts[partsIndexCounter], 32)
		if err != nil {
			return values, err
		}
		// todo possible data check
		values[valuesIndexCounter] = parsedValue
	}
	return values, nil

}

func parseIntConvertToUint(toParse string) (uint, error) {
	result, err := strconv.Atoi(toParse)
	if err != nil {
		return 0, err
	}
	if result < 0 {
		return 0, errors.New("parsed value is below zero")
	}
	return uint(result), nil
}

func parsePacketTime(packetParts []string) (uint, error) {
	time, err := parseIntConvertToUint(packetParts[parsedDataPacketIndexes.time])
	if err != nil {
		return 0, err
	}
	return time, nil
}

func parsePacketNumber(packetParts []string) (uint, error) {
	packetNumber, err := parseIntConvertToUint(packetParts[parsedDataPacketIndexes.packetNumber])
	if err != nil {
		return 0, err
	}
	return packetNumber, nil
}

func parsePacketDeviceID(packetParts []string) (uint, error) {
	packetNumber, err := parseIntConvertToUint(packetParts[parsedDataPacketIndexes.deviceID])
	if err != nil {
		return 0, errors.New("failed to parse packet number")
	}
	return packetNumber, nil
}

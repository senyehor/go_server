package packet

import (
	"errors"
	"github.com/senyehor/go_server/utils"
	"strconv"
	"strings"
)

var (
	packetConfig = utils.GetPacketConfig()
)

func getPacketPartsIndexesInParsedData() *packetPartsIndexesInParsedData {
	// [Token];[n1];[n2];...;[packetConfig.ValuesCount()];[Time];[PacketNumber];[IDdevice]!
	//- Packet structure
	return &packetPartsIndexesInParsedData{
		token:              0,
		valuesRangeBorders: newRangeBorders(1, packetConfig.ValuesCount()+1), // left border included, second excluded
		// indexes below are dependent on ValuesCount
		//and each shifts to one more from right border of values right border
		time:         1 + packetConfig.ValuesCount(),
		packetNumber: 2 + packetConfig.ValuesCount(),
		deviceID:     3 + packetConfig.ValuesCount(),
	}
}

func parseBinaryDataToStringParts(binaryData []byte) ([]string, error) {
	parts := strings.Split(string(binaryData[:]), string(packetConfig.DataDelimiter()))
	if !checkPacketLength(parts) {
		return parts, errors.New("invalid Packet length")
	}
	if !checkPacketToken(parts) {
		return parts, errors.New("invalid Packet token")
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
	return uint8(len(packetParts)) == packetConfig.NonValuesPartsCount()+packetConfig.ValuesCount()
}

func checkPacketToken(packetParts []string) bool {
	return strings.Compare(packetParts[parsedDataPacketIndexes.token], packetConfig.Token()) == 0
}

func parsePacketValues(packetParts []string) (packetValues, error) {
	var values packetValues
	for partsIndexCounter := parsedDataPacketIndexes.valuesRangeBorders.left; partsIndexCounter < parsedDataPacketIndexes.valuesRangeBorders.right; partsIndexCounter++ {
		parsedValue, err := strconv.ParseFloat(packetParts[partsIndexCounter], 32)
		if err != nil {
			return values, err
		}
		values = append(values, parsedValue)
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
		return 0, errors.New("failed to parse Packet number")
	}
	return packetNumber, nil
}

package binary_parser

import (
	"errors"
	"strconv"
	"strings"
)

func parseBinaryDataToStringParts(binaryData []byte) (*incomingDataStringParts, error) {
	rawPacketParts := strings.Split(string(binaryData[:]), string(packetConfig.DataDelimiter()))
	if !checkPacketLength(rawPacketParts) {
		return nil, errors.New("invalid packet length")
	}
	trimTerminator(rawPacketParts)
	return newIncomingDataPartsFromArray(rawPacketParts), nil
}

func checkPacketLength(packetParts []string) bool {
	return len(packetParts) == packetConfig.OtherValuesCount()+packetConfig.ValuesCount()
}

func trimTerminator(packetParts []string) {
	lastElementPosition := len(packetParts) - 1
	packetParts[lastElementPosition] = strings.TrimRight(
		packetParts[lastElementPosition],
		string(packetConfig.DataTerminator()),
	)
	// catching index out of range err
	if err := recover(); err != nil {
		return
	}
}

func checkPacketToken(token string) bool {
	return strings.Compare(token, packetConfig.Token()) == 0
}

func parsePacketValues(incomingValuesToParse []string) ([]float64, error) {
	values := make([]float64, len(incomingValuesToParse))
	for index, _ := range values {
		parsedValue, err := strconv.ParseFloat(incomingValuesToParse[index], 64)
		if err != nil {
			return nil, errors.New("failed to parse a packet value")
		}
		values[index] = parsedValue
	}
	return values, nil
}

func parsePacketTimeInterval(packetTimeIntervalToParse string) (int, error) {
	timeInterval, err := strconv.ParseInt(packetTimeIntervalToParse, 10, 32)
	if err != nil {
		return 0, errors.New("failed to parse packet timeInterval")
	}
	return int(timeInterval), nil
}

func parsePacketNumber(packetNumberToParse string) (int, error) {
	packetNumber, err := strconv.ParseInt(packetNumberToParse, 10, 32)
	if err != nil {
		return 0, errors.New("failed to parse packet number")
	}
	return int(packetNumber), nil
}

func parsePacketDeviceID(packetDeviceIDToParse string) (int, error) {
	packetNumber, err := strconv.ParseInt(packetDeviceIDToParse, 10, 32)
	if err != nil {
		return 0, errors.New("failed to parse packet device id")
	}
	return int(packetNumber), nil
}

package binary_parser

import (
	"errors"
	"strconv"
	"strings"
)

func parseBinaryDataToStringParts(binaryData []byte) (*incomingDataStringParts, error) {
	arrayOfParsedItems := strings.Split(string(binaryData[:]), string(packetConfig.DataDelimiter()))
	if !checkPacketLength(arrayOfParsedItems) {
		return nil, errors.New("invalid packet length")
	}

	parts := newIncomingDataStringPartsFromArray(arrayOfParsedItems)
	if !checkPacketToken(parts.Token()) {
		return nil, errors.New("invalid packet token")
	}
	return parts, nil
}

func checkPacketLength(packetParts []string) bool {
	return len(packetParts) == packetConfig.OtherValuesCount()+packetConfig.ValuesCount()
}

func checkPacketToken(token string) bool {
	return strings.Compare(token, packetConfig.Token()) == 0
}

func parsePacketValues(incomingValuesToParse []string) ([]float64, error) {
	values := make([]float64, packetConfig.ValuesCount(), packetConfig.ValuesCount())
	for partsIndexCounter, _ := range values {
		parsedValue, err := strconv.ParseFloat(incomingValuesToParse[partsIndexCounter], 64)
		if err != nil {
			return nil, errors.New("failed to parse a packet value")
		}
		values = append(values, parsedValue)
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

// getPacketPartsIndexesInParsedData
// [Token];[n1];[n2];...;[packetConfig.ValuesCount()];[TimeInterval];[PacketNumber];[IDdevice]!
//- Packet structure
func getPacketPartsIndexesInParsedData() *packetPartsIndexesInParsedData {
	return &packetPartsIndexesInParsedData{
		token: 0,
		// left border included, right excluded
		valuesLeftBorder:  1,
		valuesRightBorder: 1 + packetConfig.ValuesCount(),
		// indexes below are dependent on ValuesCount
		//and each shifts to one more from values right border
		time:         1 + packetConfig.ValuesCount(),
		packetNumber: 2 + packetConfig.ValuesCount(),
		deviceID:     3 + packetConfig.ValuesCount(),
	}
}

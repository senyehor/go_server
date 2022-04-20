package packet

import (
	"errors"
	"github.com/senyehor/go_server/utils"
	"strconv"
	"strings"
)

var (
	packetConfig            = utils.GetPacketConfig()
	parsedDataPacketIndexes = getPacketPartsIndexesInParsedData()
)

// getPacketPartsIndexesInParsedData
// [Token];[n1];[n2];...;[packetConfig.ValuesCount()];[TimeInterval];[PacketNumber];[IDdevice]!
//- Packet structure
func getPacketPartsIndexesInParsedData() *packetPartsIndexesInParsedData {
	return &packetPartsIndexesInParsedData{
		token: 0,
		// left border included, second excluded
		valuesRangeBorders: newRangeBorders(1, packetConfig.ValuesCount()+1),
		// indexes below are dependent on ValuesCount
		//and each shifts to one more from values right border
		time:         1 + packetConfig.ValuesCount(),
		packetNumber: 2 + packetConfig.ValuesCount(),
		deviceID:     3 + packetConfig.ValuesCount(),
	}
}

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
	return uint(len(packetParts)) == packetConfig.OtherValuesCount()+packetConfig.ValuesCount()
}

func checkPacketToken(token string) bool {
	return strings.Compare(token, packetConfig.Token()) == 0
}

func parsePacketValues(incomingValuesToParse []string) (*packetValues, error) {
	values := newEmptyPacketValues()
	for partsIndexCounter := uint(0); partsIndexCounter < packetConfig.ValuesCount(); partsIndexCounter++ {
		parsedValue, err := strconv.ParseFloat(incomingValuesToParse[partsIndexCounter], 64)
		if err != nil {
			return nil, errors.New("failed to parse a packet value")
		}
		if !values.Append(parsedValue) {
			return nil, errors.New("exceeded values capacity")
		}
	}
	return values, nil
}

func parsePacketTimeInterval(packetTimeIntervalToParse string) (uint, error) {
	timeInterval, err := utils.ParseIntConvertToUint(packetTimeIntervalToParse)
	if err != nil {
		return 0, errors.New("failed to parse packet timeInterval")
	}
	return timeInterval, nil
}

func parsePacketNumber(packetNumberToParse string) (uint, error) {
	packetNumber, err := utils.ParseIntConvertToUint(packetNumberToParse)
	if err != nil {
		return 0, errors.New("failed to parse packet number")
	}
	return packetNumber, nil
}

func parsePacketDeviceID(packetDeviceIDToParse string) (uint, error) {
	packetNumber, err := utils.ParseIntConvertToUint(packetDeviceIDToParse)
	if err != nil {
		return 0, errors.New("failed to parse packet device id")
	}
	return packetNumber, nil
}

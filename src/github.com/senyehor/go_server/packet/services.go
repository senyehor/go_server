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

func getPacketPartsIndexesInParsedData() *packetPartsIndexesInParsedData {
	// [Token];[n1];[n2];...;[packetConfig.ValuesCount()];[TimeInterval];[PacketNumber];[IDdevice]!
	//- Packet structure
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
	parsedArray := strings.Split(string(binaryData[:]), string(packetConfig.DataDelimiter()))
	if !checkPacketLength(parsedArray) {
		return nil, errors.New("invalid packet length")
	}

	parts := newIncomingDataStringPartsFromArray(parsedArray)
	if !checkPacketToken(parts.Token()) {
		return nil, errors.New("invalid packet token")
	}
	return parts, nil
}

func checkPacketLength(packetParts []string) bool {
	return uint8(len(packetParts)) == packetConfig.NonValuesPartsCount()+packetConfig.ValuesCount()
}

func checkPacketToken(token string) bool {
	return strings.Compare(token, packetConfig.Token()) == 0
}

func parsePacketValues(incomingValuesToParse []string) (*packetValues, error) {
	values := newPacketValues()
	for partsIndexCounter := uint8(0); partsIndexCounter < packetConfig.ValuesCount(); partsIndexCounter++ {
		parsedValue, err := strconv.ParseFloat(incomingValuesToParse[partsIndexCounter], 64)
		if err != nil {
			return nil, errors.New("failed to parse a packet value")
		}
		values.Append(parsedValue)
	}
	return values, nil

}

func parsePacketTime(packetTimeToParse string) (uint, error) {
	timeInterval, err := utils.ParseIntConvertToUint(packetTimeToParse)
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

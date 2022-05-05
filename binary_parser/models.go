package binary_parser

import (
	"github.com/senyehor/go_server/utils"
)

var (
	packetConfig            = utils.PacketConfig
	parsedDataPacketIndexes = newParsedDataPacketIndexes()
)

type incomingDataStringParts struct {
	token  string
	values []string
	timeInterval,
	packetNumber,
	deviceID string
}

type packetPartsIndexesInParsedData struct {
	token,
	valuesLeftBorder,
	valuesRightBorder,
	timeInterval,
	packetNumber,
	deviceID int
}

func newIncomingDataStringParts(
	token string, values []string, time string, packetNumber string, deviceID string) *incomingDataStringParts {
	valuesCopy := values[:]
	return &incomingDataStringParts{token: token, values: valuesCopy, timeInterval: time, packetNumber: packetNumber,
		deviceID: deviceID}
}

func newIncomingDataPartsFromArray(parts []string) *incomingDataStringParts {
	return newIncomingDataStringParts(
		parts[parsedDataPacketIndexes.token],
		parts[parsedDataPacketIndexes.valuesLeftBorder:parsedDataPacketIndexes.valuesRightBorder],
		parts[parsedDataPacketIndexes.timeInterval],
		parts[parsedDataPacketIndexes.packetNumber],
		parts[parsedDataPacketIndexes.deviceID],
	)
}

func (i *incomingDataStringParts) Copy() *incomingDataStringParts {
	valuesCopy := make([]string, packetConfig.ValuesCount())
	copy(valuesCopy, i.Values())
	return &incomingDataStringParts{
		values:       valuesCopy,
		token:        i.Token(),
		timeInterval: i.TimeInterval(),
		packetNumber: i.PacketNumber(),
		deviceID:     i.DeviceID(),
	}
}
func (i *incomingDataStringParts) Values() []string {
	return i.values
}
func (i *incomingDataStringParts) Token() string {
	return i.token
}
func (i *incomingDataStringParts) TimeInterval() string {
	return i.timeInterval
}
func (i *incomingDataStringParts) PacketNumber() string {
	return i.packetNumber
}
func (i *incomingDataStringParts) DeviceID() string {
	return i.deviceID
}
func (i *incomingDataStringParts) Equal(other *incomingDataStringParts) bool {
	if i.TimeInterval() != other.TimeInterval() {
		return false
	}
	if i.PacketNumber() != other.PacketNumber() {
		return false
	}
	if i.DeviceID() != other.DeviceID() {
		return false
	}
	if len(i.Values()) != len(other.Values()) {
		return false
	}
	for index, value := range i.Values() {
		if other.Values()[index] != value {
			return false
		}
	}
	return true
}

// [Token];[n1];[n2];...;[packetConfig.ValuesCount()];[TimeInterval];[PacketNumber];[IDdevice]! - packet structure
func newParsedDataPacketIndexes() *packetPartsIndexesInParsedData {
	return &packetPartsIndexesInParsedData{
		token: 0,
		// left border included, right excluded
		valuesLeftBorder:  1,
		valuesRightBorder: 1 + packetConfig.ValuesCount(),
		// indexes below are dependent on ValuesCount
		//and each shifts to one more from values right border
		timeInterval: 1 + packetConfig.ValuesCount(),
		packetNumber: 2 + packetConfig.ValuesCount(),
		deviceID:     3 + packetConfig.ValuesCount(),
	}
}

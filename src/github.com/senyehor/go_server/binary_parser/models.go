package binary_parser

import (
	"github.com/senyehor/go_server/utils"
	"strings"
)

var (
	packetConfig            = utils.GetPacketConfig()
	parsedDataPacketIndexes = getPacketPartsIndexesInParsedData()
)

type incomingDataStringParts struct {
	token  string
	values []string
	time,
	packetNumber,
	deviceID string
}

type packetPartsIndexesInParsedData struct {
	token,
	valuesLeftBorder,
	valuesRightBorder,
	time,
	packetNumber,
	deviceID int
}

func newIncomingDataStringParts(
	token string, values []string, time string, packetNumber string, deviceID string) *incomingDataStringParts {
	return &incomingDataStringParts{token: token, values: values, time: time, packetNumber: packetNumber,
		deviceID: deviceID}
}
func newIncomingDataStringPartsFromArray(parts []string) *incomingDataStringParts {
	valuesCopy := make([]string, packetConfig.ValuesCount())
	copy(
		valuesCopy,
		parts[parsedDataPacketIndexes.valuesLeftBorder:parsedDataPacketIndexes.valuesLeftBorder],
	)
	trimmedFromTerminatorDeviceID := strings.TrimRight(
		parts[parsedDataPacketIndexes.deviceID],
		string(packetConfig.DataTerminator()),
	)
	return newIncomingDataStringParts(
		parts[parsedDataPacketIndexes.token],
		valuesCopy,
		parts[parsedDataPacketIndexes.time],
		parts[parsedDataPacketIndexes.packetNumber],
		trimmedFromTerminatorDeviceID,
	)
}

func (i *incomingDataStringParts) IsEqual(other *incomingDataStringParts) bool {
	if i.Token() != other.Token() {
		return false
	}
	if len(i.Values()) != len(other.Values()) {
		return false
	}
	for index, value := range i.values {
		if value != other.values[index] {
			return false
		}
	}
	if i.Time() != other.Time() {
		return false
	}
	if i.PacketNumber() != other.PacketNumber() {
		return false
	}
	return i.DeviceID() == other.DeviceID()
}
func (i *incomingDataStringParts) Copy() *incomingDataStringParts {
	valuesCopy := make([]string, packetConfig.ValuesCount())
	copy(valuesCopy, i.Values())
	return &incomingDataStringParts{
		values:       valuesCopy,
		token:        i.Token(),
		time:         i.Time(),
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
func (i *incomingDataStringParts) Time() string {
	return i.time
}
func (i *incomingDataStringParts) PacketNumber() string {
	return i.packetNumber
}
func (i *incomingDataStringParts) DeviceID() string {
	return i.deviceID
}

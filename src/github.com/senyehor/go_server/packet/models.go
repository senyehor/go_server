package packet

import (
	"strings"
)

type packetValues struct {
	values []float64
}

type packetValuesIterationReturn struct {
	isLast        bool
	value         float64
	valuePosition int
}

type Packet struct {
	values       *packetValues
	timeInterval uint
	packetNum    uint
	deviceID     uint
}

type rangeBorders struct {
	left  uint
	right uint
}

type packetPartsIndexesInParsedData struct {
	token              uint
	valuesRangeBorders *rangeBorders
	time               uint
	packetNumber       uint
	deviceID           uint
}

type incomingDataStringParts struct {
	token  string
	values []string
	time,
	packetNumber,
	deviceID string
}

func newEmptyPacketValues() *packetValues {
	return &packetValues{values: []float64{}}
}
func (p *packetValues) Append(value float64) bool {
	if uint(len(p.values)) == packetConfig.ValuesCount() {
		return false
	}
	p.values = append(p.values, value)
	return true
}

// Iterate returns if element is last, element itself, and it`s position
func (p *packetValues) Iterate() <-chan *packetValuesIterationReturn {

	channel := make(chan *packetValuesIterationReturn)
	go func() {
		length := len(p.values)
		for index, value := range p.values {
			channel <- newPacketValuesIterationReturn(index == length-1, value, index)
		}
		close(channel)
	}()
	return channel
}

func newPacketValuesIterationReturn(
	isLast bool, value float64, valuePosition int) *packetValuesIterationReturn {
	return &packetValuesIterationReturn{isLast: isLast, value: value, valuePosition: valuePosition}
}
func (p *packetValuesIterationReturn) IsLast() bool {
	return p.isLast
}
func (p *packetValuesIterationReturn) Value() float64 {
	return p.value
}
func (p *packetValuesIterationReturn) ValuePosition() int {
	return p.valuePosition
}

func NewPacket(values *packetValues, timeInterval uint, packetNum uint, deviceID uint) *Packet {
	return &Packet{values: values, timeInterval: timeInterval, packetNum: packetNum, deviceID: deviceID}
}
func (p *Packet) Values() *packetValues {
	return p.values
}
func (p *Packet) TimeInterval() uint {
	return p.timeInterval
}
func (p *Packet) PacketNum() uint {
	return p.packetNum
}
func (p *Packet) DeviceID() uint {
	return p.deviceID
}

func newRangeBorders(left uint, right uint) *rangeBorders {
	return &rangeBorders{left: left, right: right}
}
func (rangeBorders *rangeBorders) Left() uint {
	return rangeBorders.left
}
func (rangeBorders *rangeBorders) Right() uint {
	return rangeBorders.right
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
		parts[parsedDataPacketIndexes.valuesRangeBorders.left:parsedDataPacketIndexes.valuesRangeBorders.right],
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

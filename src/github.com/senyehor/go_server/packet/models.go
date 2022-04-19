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
	left  uint8
	right uint8
}

type packetPartsIndexesInParsedData struct {
	token              uint8
	valuesRangeBorders *rangeBorders
	time               uint8
	packetNumber       uint8
	deviceID           uint8
}

type incomingDataStringParts struct {
	token  string
	values []string
	time,
	packetNumber,
	deviceID string
}

func newPacketValues() *packetValues {
	return &packetValues{values: []float64{}}
}
func (p *packetValues) Append(value float64) {
	if uint8(len(p.values)) == packetConfig.ValuesCount() {
		// todo remake to error
		panic("exceeded values part capacity")
	}
	p.values = append(p.values, value)
}
func (p *packetValues) Iterate() <-chan *packetValuesIterationReturn {
	// returns if element is last, element itself, and it`s position
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

func NewPacket(values *packetValues, time uint, packetNum uint, deviceID uint) *Packet {
	return &Packet{values: values, timeInterval: time, packetNum: packetNum, deviceID: deviceID}
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

func newRangeBorders(left uint8, right uint8) *rangeBorders {
	return &rangeBorders{left: left, right: right}
}
func (rangeBorders *rangeBorders) Left() uint8 {
	return rangeBorders.left
}
func (rangeBorders *rangeBorders) Right() uint8 {
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

package packet

type packetValues struct {
	items []float64
}
type packetValuesIterationReturn struct {
	isLast        bool
	value         float64
	valuePosition uint
}

func (p packetValues) Iterate() <-chan packetValuesIterationReturn {
	// returns if element is last, element itself, and it`s position
	channel := make(chan packetValuesIterationReturn)
	go func() {
		length := len(p.items)
		for index, value := range p.items {
			channel <- packetValuesIterationReturn{
				isLast:        index == length-1,
				value:         value,
				valuePosition: uint(index) + 1,
			}
		}
		close(channel)
	}()
	return channel
}

type Packet struct {
	values    packetValues
	time      uint
	packetNum uint
	deviceID  uint
}

type rangeBorders struct {
	left  uint8
	right uint8
}

func newRangeBorders(left uint8, right uint8) *rangeBorders {
	return &rangeBorders{left: left, right: right}
}

type packetPartsIndexesInParsedData struct {
	token              uint8
	valuesRangeBorders *rangeBorders
	time               uint8
	packetNumber       uint8
	deviceID           uint8
}

func (p packetValuesIterationReturn) IsLast() bool {
	return p.isLast
}
func (p packetValuesIterationReturn) Value() float64 {
	return p.value
}
func (p packetValuesIterationReturn) ValuePosition() uint {
	return p.valuePosition
}

func NewPacket(values packetValues, time uint, packetNum uint, deviceID uint) *Packet {
	return &Packet{values: values, time: time, packetNum: packetNum, deviceID: deviceID}
}
func (p Packet) Values() *packetValues {
	return &p.values
}
func (p Packet) Time() uint {
	return p.time
}
func (p Packet) PacketNum() uint {
	return p.packetNum
}
func (p Packet) DeviceID() uint {
	return p.deviceID
}

func (rangeBorders *rangeBorders) Left() uint8 {
	return rangeBorders.left
}
func (rangeBorders *rangeBorders) Right() uint8 {
	return rangeBorders.right
}

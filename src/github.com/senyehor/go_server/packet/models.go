package packet

type packetValues = []float64

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

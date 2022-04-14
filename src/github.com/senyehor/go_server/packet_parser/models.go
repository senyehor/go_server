package packet_parser

const ( // must be sync to type Packet
	PacketValuesCount         = 16
	NonValuesPacketPartsCount = 4
)

type packetValues = [PacketValuesCount]float64

type Packet struct { // must be kept to constants above
	values    packetValues
	time      uint
	packetNum uint
	deviceID  uint
}

type rangeBorders struct {
	left  uint8
	right uint8
}

type packetPartsIndexesInParsedData struct {
	token              uint8
	valuesRangeBorders rangeBorders
	time               uint8
	packetNumber       uint8
	deviceID           uint8
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

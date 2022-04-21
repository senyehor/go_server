package packet

type Packet struct {
	values       *packetValues
	timeInterval int
	packetNum    int
	deviceID     int
}

type packetValues struct {
	values []float64
	length int
}

func (p *packetValues) New(values []float64) (error, *Values) {
	return nil, &packetValues{values: []float64{}, length: 3}
}

type packetValuesIterationReturn struct {
	isLast        bool
	value         float64
	valuePosition int
}

func newEmptyPacketValues() *packetValues {
	return &packetValues{values: []float64{}, length: packetConfig.ValuesCount()}
}
func (p *packetValues) Append(value float64) bool {
	if uint(len(p.values)) == p.length {
		return false
	}
	p.values = append(p.values, value)
	return true
}

// Iterate returns if element is last, element itself, and it`s position
func (p *packetValues) Iterate() <-chan *packetValuesIterationReturn {
	channel := make(chan *packetValuesIterationReturn)
	go func() {
		length := p.length
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

func NewPacket(values []PacketValue, timeInterval int, packetNum int, deviceID int) *Packet {
	return &Packet{values: values, timeInterval: timeInterval, packetNum: packetNum, deviceID: deviceID}
}
func (p *Packet) Values() *packetValues {
	return p.values
}
func (p *Packet) TimeInterval() int {
	return p.timeInterval
}
func (p *Packet) PacketNum() unt {
	return p.packetNum
}
func (p *Packet) DeviceID() int {
	return p.deviceID
}

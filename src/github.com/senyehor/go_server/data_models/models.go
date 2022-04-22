package data_models

import (
	"errors"
	"github.com/senyehor/go_server/utils"
)

type Packet struct {
	values       *PacketValues
	timeInterval int
	packetNum    int
	deviceID     int
}

func NewPacket(values []float64, timeInterval int, packetNum int, deviceID int) (*Packet, error) {
	packetValues, err := newPacketValues(values)
	if err != nil {
		return nil, errors.New("incorrect Packet values")
	}
	if timeInterval < 0 {
		return nil, errors.New("Packet time interval below zero")
	}
	if packetNum < 0 {
		return nil, errors.New("Packet number below zero")
	}
	if timeInterval < 0 {
		return nil, errors.New("Packet device id below zero")
	}
	return &Packet{values: packetValues, timeInterval: timeInterval, packetNum: packetNum, deviceID: deviceID}, nil
}
func (p *Packet) Values() *PacketValues {
	return p.values
}
func (p *Packet) TimeInterval() int {
	return p.timeInterval
}
func (p *Packet) PacketNum() int {
	return p.packetNum
}
func (p *Packet) DeviceID() int {
	return p.deviceID
}

type PacketValues struct {
	values []float64
}

func newPacketValues(values []float64) (*PacketValues, error) {
	valuesCount := utils.PacketConfig.ValuesCount()
	if len(values) != valuesCount {
		return nil, errors.New("passed PacketValues count did not match expected")
	}
	valuesCopy := make([]float64, valuesCount, valuesCount)
	copy(valuesCopy, values)
	return &PacketValues{values: valuesCopy}, nil
}
func (pvs *PacketValues) Iterator() *PacketValuesIterator {
	return newPacketValuesIterator(pvs)
}

type PacketValuesIterator struct {
	packetValues         *PacketValues
	valuesCount          int
	currentValue         float64
	currentValuePosition int
}

func newPacketValuesIterator(packetValues *PacketValues) *PacketValuesIterator {
	return &PacketValuesIterator{packetValues: packetValues, valuesCount: len(packetValues.values)}
}
func (pvi *PacketValuesIterator) HasNext() bool {
	if pvi.currentValuePosition == pvi.valuesCount {
		return false
	}
	if pvi.currentValuePosition == 0 {
		pvi.currentValue = pvi.packetValues.values[pvi.currentValuePosition]
		return true
	}
	pvi.currentValuePosition++
	return true
}
func (pvi *PacketValuesIterator) Value() float64 {
	return pvi.currentValue
}
func (pvi *PacketValuesIterator) ValuePosition() int {
	return pvi.currentValuePosition
}
func (pvi *PacketValuesIterator) IsLast() bool {
	return pvi.currentValuePosition == pvi.valuesCount-1
}

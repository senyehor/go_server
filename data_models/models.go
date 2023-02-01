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
		return nil, errors.New("incorrect packet values")
	}
	if timeInterval < 0 {
		return nil, errors.New("packet time interval below zero")
	}
	if packetNum < 0 {
		return nil, errors.New("packet number below zero")
	}
	if deviceID < 0 {
		return nil, errors.New("packet device id below zero")
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
	valuesInKilowattHour := convertJouleValuesToKilowattHour(values)
	return &PacketValues{values: valuesInKilowattHour}, nil
}

//Iterator returns an iterator that should be iterated by
//
//for iterator.HasNext() {here you access results of iteration via iterator.Something()}
func (pvs *PacketValues) Iterator() *PacketValuesIterator {
	return newPacketValuesIterator(pvs)
}

type PacketValuesIterator struct {
	packetValues     *PacketValues
	length           int
	currentItem      float64
	iterationCounter int
}

func newPacketValuesIterator(packetValues *PacketValues) *PacketValuesIterator {
	return &PacketValuesIterator{
		packetValues:     packetValues,
		length:           len(packetValues.values),
		currentItem:      0,
		iterationCounter: 0,
	}
}
func (pvi *PacketValuesIterator) HasNext() bool {
	if pvi.iterationCounter == pvi.length {
		return false
	}
	pvi.currentItem = pvi.packetValues.values[pvi.iterationCounter]
	pvi.iterationCounter++
	return true
}
func (pvi *PacketValuesIterator) Value() float64 {
	pvi.checkIterationStarted()
	return pvi.currentItem
}

func (pvi *PacketValuesIterator) checkIterationStarted() {
	if pvi.iterationCounter == 0 {
		// todo think of better solution
		panic("HasNext was not called")
	}
}

// ValuePosition return value position from 0
func (pvi *PacketValuesIterator) ValuePosition() int {
	pvi.checkIterationStarted()
	return pvi.iterationCounter - 1
}
func (pvi *PacketValuesIterator) IsLast() bool {
	pvi.checkIterationStarted()
	return pvi.iterationCounter == pvi.length
}

package data_models

import (
	"github.com/senyehor/go_server/utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestDataModels(t *testing.T) {
	suite.Run(t, new(appModelsTestSuite))
}

type appModelsTestSuite struct {
	suite.Suite
	correctValues      []float64
	incorrectValuesSet [][]float64
	correctTimeInterval,
	correctPacketNumber,
	correctDeviceID int
}

func (s *appModelsTestSuite) SetupTest() {
	// Setup set is called multiple times, so there is no point to reinitialize values each time
	if s.correctValues == nil {
		// filling with random negative and positive numbers and ensuring values has correct length
		for i := 0; i < utils.PacketConfig.ValuesCount(); i++ {
			s.correctValues = append(s.correctValues, utils.RandFloat64())
		}
		// creating possible incorrect values
		longerValues := make([]float64, utils.PacketConfig.ValuesCount())
		copy(longerValues, s.correctValues)
		longerValues = append(longerValues, utils.RandFloat64())
		s.incorrectValuesSet = append(s.incorrectValuesSet, longerValues)
		shorterValues := longerValues[:len(longerValues)-2]
		s.incorrectValuesSet = append(s.incorrectValuesSet, shorterValues)
		s.incorrectValuesSet = append(s.incorrectValuesSet, nil)
	}

	s.correctTimeInterval = utils.RandPositiveInt()
	s.correctPacketNumber = utils.RandPositiveInt()
	s.correctDeviceID = utils.RandPositiveInt()
}

func (s *appModelsTestSuite) TestCreatingPacket() {
	result, err := NewPacket(s.correctValues, s.correctTimeInterval, s.correctPacketNumber, s.correctDeviceID)
	s.NoError(err, "NewPacket returned err with correct incoming data")
	valuesIterator := result.Values().Iterator()
	for valuesIterator.HasNext() {
		s.True(
			utils.CompareFloatsPrecise(valuesIterator.Value(), s.correctValues[valuesIterator.ValuePosition()]),
			"packet values does not match expected",
		)
	}
	s.Equal(s.correctTimeInterval, result.TimeInterval(), "packet was created with wrong time interval")
	s.Equal(s.correctPacketNumber, result.PacketNum(), "packet was created with wrong number")
	s.Equal(s.correctDeviceID, result.DeviceID(), "packet was created with wrong device id")

	for _, incorrectValues := range s.incorrectValuesSet {
		result, err := NewPacket(incorrectValues, s.correctTimeInterval, s.correctPacketNumber, s.correctDeviceID)
		s.Nil(result, "NewPacket did not return nil, when wrong values were passed")
		s.Error(err, "newPacketValues did not return error, when wrong values were passed")
	}

	// currently, only limitation for time interval, packet number is device id is to be > 0
	result, err = NewPacket(s.correctValues, s.correctTimeInterval*(-1.0), s.correctPacketNumber, s.correctDeviceID)
	s.Nil(result, "NewPacket did not return nil, when negative time interval was passed")
	s.Error(err, "newPacketValues did not return error, when negative time interval was passed")

	result, err = NewPacket(s.correctValues, s.correctTimeInterval, s.correctPacketNumber*(-1.0), s.correctDeviceID)
	s.Nil(result, "NewPacket did not return nil, when negative packet number was passed")
	s.Error(err, "newPacketValues did not return error, when negative packet number was passed")

	result, err = NewPacket(s.correctValues, s.correctTimeInterval, s.correctPacketNumber, s.correctDeviceID*(-1.0))
	s.Nil(result, "NewPacket did not return nil, when negative packet number was passed")
	s.Error(err, "newPacketValues did not return error, when negative packet number was passed")
}

func (s *appModelsTestSuite) TestCreatingPacketValues() {
	result, err := newPacketValues(s.correctValues)
	s.NoError(err, "packet values were not created with correct incoming data")
	s.Equal(len(s.correctValues), len(result.values), "packet values were made with incorrect length")
	for index, value := range result.values {
		s.True(
			utils.CompareFloatsPrecise(s.correctValues[index], value),
			"packet values incorrectValuesSet did not match expected",
		)
	}

	for _, incorrectValues := range s.incorrectValuesSet {
		result, err = newPacketValues(incorrectValues)
		s.Nil(result, "newPacketValues did not return nil, when wrong data was passed")
		s.Error(err, "newPacketValues did not return error, when wrong data was passed")
	}
}

func (s *appModelsTestSuite) TestValueIterator() {
	packetValues, err := newPacketValues(s.correctValues)
	if err != nil {
		s.Fail("newPacketValues returned err while provided correct values")
	}

	iterator := packetValues.Iterator()
	for range s.correctValues {
		s.True(iterator.HasNext(), "values iterator has less iterations than expected")
	}
	s.False(iterator.HasNext(), "values iterator has more iterations than expected")

	// in this part we presume iterator HasNext method works correctly
	iterator = packetValues.Iterator()
	for index, value := range s.correctValues {
		iterator.HasNext() // calling method to iterate
		s.Equal(index, iterator.ValuePosition(), "iterator ValuePosition did not work as expected")
		s.True(
			utils.CompareFloatsPrecise(value, iterator.Value()),
			"iterator Value did not work as expected",
		)
	}
}

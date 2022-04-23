package binary_parser

import (
	"fmt"
	"github.com/senyehor/go_server/utils"
	"github.com/stretchr/testify/suite"
	"strconv"
	"strings"
	"testing"
)

func TestBinaryParser(t *testing.T) {
	suite.Run(t, new(binaryParserTestSuite))
}

type binaryParserTestSuite struct {
	suite.Suite
	correctParts               *incomingDataStringParts
	correctPartsExpectedValues *correctPartsExpectedValues
}

type correctPartsExpectedValues struct {
	token  string
	values []float64
	timeInterval,
	packetNumber,
	deviceID int
}

func (s *binaryParserTestSuite) SetupTest() {
	var stringValues []string
	var values []float64
	for i := 0; i < packetConfig.ValuesCount(); i++ {
		randFloat := utils.RandFloat64()
		values = append(values, randFloat)
		stringValues = append(stringValues, fmt.Sprintf("%v", randFloat))
	}
	timeInterval := utils.RandPositiveInt()
	packetNumber := utils.RandPositiveInt()
	deviceID := utils.RandPositiveInt()
	s.correctPartsExpectedValues = &correctPartsExpectedValues{
		token:        packetConfig.Token(),
		values:       values,
		timeInterval: timeInterval,
		packetNumber: packetNumber,
		deviceID:     deviceID,
	}
	s.correctParts = newIncomingDataStringParts(
		packetConfig.Token(),
		stringValues,
		strconv.Itoa(timeInterval),
		strconv.Itoa(packetNumber),
		strconv.Itoa(deviceID),
	)
}

func (s *binaryParserTestSuite) TestParseFromBinary() {
	result, err := ParseFromBinary(s.composeCorrectBinaryData())
	s.NoError(err, "ParseFromBinary returned err with correct input")
	correctValues := s.correctPartsExpectedValues
	s.Equal(result.DeviceID(), correctValues.deviceID, "device id did not match expected")
	s.Equal(result.TimeInterval(), correctValues.timeInterval, "time interval did not match expected")
	s.Equal(result.PacketNum(), correctValues.packetNumber, "packet number did not match expected")
	iterator := result.Values().Iterator()
	valuesLength := 0
	for iterator.HasNext() {
		valuesLength++
	}
	s.Len(correctValues.values, valuesLength, "incorrect packet values length")
	iterator = result.Values().Iterator()
	for iterator.HasNext() {
		s.True(utils.CompareFloatsPrecise(iterator.Value(),
			correctValues.values[iterator.ValuePosition()]),
			"packet value did not match expected",
		)
	}

	randomData := newIncomingDataStringParts("random token",
		[]string{"some value 1", "some value 2"},
		"incorrect time",
		"-124234",
		"$#*&^",
	)

	incorrectInputs := [][]byte{
		[]byte(""),
		[]byte("possible random data"),
		s.composeBinaryData(randomData, "random delimiter", "random terminator"),
	}
	for _, elem := range incorrectInputs {
		result, err = ParseFromBinary(elem)
		s.Error(err, "ParseFromBinary did not return err with incorrect input")
		s.Nil(result, "ParseFromBinary did not return nil with incorrect input")
	}
}

func (s *binaryParserTestSuite) TestParseBinaryDataToStringParts() {
	result, err := parseBinaryDataToStringParts(s.composeCorrectBinaryData())
	s.NoError(err, "parseBinaryDataToStringParts returned err with correct input")
	s.True(s.correctParts.Equal(result), "parseBinaryDataToStringParts return did not match expected")

	wrongDelimiterData := s.composeBinaryData(
		s.correctParts,
		"incorrect delimiter",
		string(packetConfig.DataTerminator()),
	)
	result, err = parseBinaryDataToStringParts(wrongDelimiterData)
	s.Nil(result, "parseBinaryDataToStringParts did not return nil with incorrect input")
	s.Error(err, "parseBinaryDataToStringParts did not return err with incorrect input")
	// case with incorrect data terminator is not considered, as part, containing it will fail parsing
}

func (s *binaryParserTestSuite) TestCheckPacketToken() {
	s.True(checkPacketToken(s.correctParts.Token()), "checkPacketToken returned false for correct token")

	partsWithIncorrectToken := s.correctParts.Copy()
	partsWithIncorrectToken.token = "incorrect token"
	s.False(checkPacketToken(
		partsWithIncorrectToken.Token()),
		"checkPacketToken returned true for wrong token",
	)
}

func (s *binaryParserTestSuite) TestParsePacketTimeInterval() {
	result, err := parsePacketTimeInterval(s.correctParts.TimeInterval())
	s.NoError(err, "parsePacketTimeInterval returned err with correct input")
	s.Equal(
		s.correctPartsExpectedValues.timeInterval,
		result,
		"parsePacketTimeInterval parsed time interval incorrectly",
	)

	incorrectTimeInterval := "incorrect time interval"
	result, err = parsePacketTimeInterval(incorrectTimeInterval)
	s.Error(err, "parsePacketTimeInterval did not return err with incorrect input")
	s.Zerof(result, "parsePacketTimeInterval did not return default value on fail")
}

func (s *binaryParserTestSuite) TestParsePacketNumber() {
	result, err := parsePacketNumber(s.correctParts.PacketNumber())
	s.NoError(err, "parsePacketNumber returned err with correct input")
	s.Equal(
		s.correctPartsExpectedValues.packetNumber,
		result,
		"parsePacketNumber parsed packet number incorrectly",
	)

	incorrectPacketNumber := "incorrect packet number"
	result, err = parsePacketNumber(incorrectPacketNumber)
	s.Error(err, "parsePacketNumber did not return err with incorrect input")
	s.Zero(result, "parsePacketNumber did not return default value on fail")
}

func (s *binaryParserTestSuite) TestParsePacketDeviceID() {
	result, err := parsePacketDeviceID(s.correctParts.DeviceID())
	s.NoError(err, "parsePacketDeviceID returned err with correct input")
	s.Equal(
		s.correctPartsExpectedValues.deviceID,
		result,
		"parsePacketDeviceID parsed device ID incorrectly",
	)

	incorrectPacketDeviceID := "incorrect device id"
	result, err = parsePacketDeviceID(incorrectPacketDeviceID)
	s.Error(err, "parsePacketDeviceID did not return err with incorrect input")
	s.Zero(result, "parsePacketDeviceID did not return default value on fail")
}

func (s *binaryParserTestSuite) TestCheckPacketLength() {
	partsWithWrongLength := s.correctParts.Copy()
	partsWithWrongLength.values = append(partsWithWrongLength.values, fmt.Sprintf("%v", utils.RandFloat64()))
	wrongLengthData := s.composeBinaryData(
		partsWithWrongLength,
		string(packetConfig.DataDelimiter()),
		string(packetConfig.DataTerminator()),
	)
	result, err := parseBinaryDataToStringParts(wrongLengthData)
	s.Nil(result, "parseBinaryDataToStringParts did not return nil with incorrect length data")
	s.Error(err, "parseBinaryDataToStringParts did not return err with incorrect length data")
}

func (s *binaryParserTestSuite) TestParsePacketValues() {
	result, err := parsePacketValues(s.correctParts.Values())
	s.NoError(err, "parsePacketValues returned err with correct data")
	s.Len(result, len(s.correctParts.Values()), "parsePacketValues returned wrong length result")
	for index, expectedValue := range s.correctPartsExpectedValues.values {
		s.Require().NoError(err, "parsing generated correct values went wrong")
		s.True(
			utils.CompareFloatsPrecise(result[index], expectedValue),
			"parsePacketValues result value did not match expected")
	}
}

func (s *binaryParserTestSuite) composeCorrectBinaryData() []byte {
	parts := s.correctParts
	delimiter := string(packetConfig.DataDelimiter())
	terminator := string(packetConfig.DataTerminator())
	return []byte(
		parts.Token() + delimiter +
			strings.Join(parts.Values(), delimiter) + delimiter +
			parts.TimeInterval() + delimiter +
			parts.PacketNumber() + delimiter +
			parts.DeviceID() + terminator)
}

func (s *binaryParserTestSuite) composeBinaryData(parts *incomingDataStringParts, delimiter, terminator string) []byte {
	return []byte(
		parts.Token() + delimiter +
			strings.Join(parts.Values(), delimiter) + delimiter +
			parts.TimeInterval() + delimiter +
			parts.PacketNumber() + delimiter +
			parts.DeviceID() + terminator)
}

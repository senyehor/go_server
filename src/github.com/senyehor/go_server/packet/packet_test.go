package packet

import (
	"errors"
	"github.com/senyehor/go_server/utils"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

func TestPacket(t *testing.T) {
	suite.Run(t, new(packetTestSuite))
}

type packetTestSuite struct {
	suite.Suite
	correctIncomingDataParts *incomingDataStringParts
	expectedParsedValues     []float64
	correctDelimiter,
	correctTerminator rune
}

func (s *packetTestSuite) SetupTest() {
	s.expectedParsedValues = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4.4, 5, 6.66666}
	s.correctIncomingDataParts = s.getCorrectIncomingDataParts()
	s.correctDelimiter = packetConfig.DataDelimiter()
	s.correctTerminator = packetConfig.DataTerminator()
}

func (s *packetTestSuite) TestParsingIncomingDataWithCorrectValues() {
	correctBinaryData := s.composeIncomingBinaryData(
		s.correctIncomingDataParts,
		string(s.correctDelimiter),
		string(s.correctTerminator),
	)
	parts, err := s.testParsingBinaryDataToStringParts(correctBinaryData)
	s.NoError(err, "parsing binary data to string parts went wrong")

	s.NoError(s.testParsePacketValues(parts), "parsing packet values went wrong")
	s.NoError(s.testParsePacketTimeInterval(parts), "parsing packet time interval went wrong")
	s.NoError(s.testParsePacketNumber(parts), "parsing packet number went wrong")
	s.NoError(s.testParsePacketDeviceId(parts), "parsing packet device id went wrong")
}

func (s *packetTestSuite) TestParsingToStringPartsWithWrongDelimiter() {
	binaryDataWithWrongDelimiter := s.composeIncomingBinaryData(
		s.correctIncomingDataParts,
		"wrong delimiter",
		string(s.correctTerminator),
	)
	_, err := s.testParsingBinaryDataToStringParts(binaryDataWithWrongDelimiter)
	s.Error(err, "parsing binary data to string parts with wrong delimiter did not return error")
}

func (s *packetTestSuite) TestParsingToStringPartsWithWrongTerminator() {
	binaryDataWithWrongTerminator := s.composeIncomingBinaryData(
		s.correctIncomingDataParts,
		string(s.correctDelimiter),
		"wrong terminator",
	)
	_, err := s.testParsingBinaryDataToStringParts(binaryDataWithWrongTerminator)
	s.Error(err, "parsing binary data to string parts with wrong  did not return error")
}

func (s *packetTestSuite) testParsingBinaryDataToStringParts(incomingData []byte) (*incomingDataStringParts, error) {
	parts, err := parseBinaryDataToStringParts(incomingData)
	if err != nil {
		return nil, errors.New("parsing incoming binary data into string parts returned err")
	}
	if !s.correctIncomingDataParts.IsEqual(parts) {
		return nil, errors.New("parsing binary data into string parts went wrong")
	}
	return parts, nil
}

func (s *packetTestSuite) testParsePacketValues(packetParts *incomingDataStringParts) error {
	values, err := parsePacketValues(packetParts.Values())
	if err != nil {
		return err
	}
	for iterItem := range values.Iterate() {
		if !utils.CompareFloats(iterItem.Value(), s.expectedParsedValues[iterItem.ValuePosition()]) {
			return errors.New("parsed value does not match expected")
		}
	}
	return nil
}

func (s *packetTestSuite) testParsePacketTimeInterval(packetParts *incomingDataStringParts) error {
	time, err := parsePacketTimeInterval(packetParts.Time())
	if err != nil {
		return errors.New("parsing packet timeInterval part returner err")
	}
	if cast.ToUint(packetParts.Time()) != time {
		return errors.New("packet time was not parsed correctly")
	}
	return nil
}

func (s *packetTestSuite) testParsePacketNumber(packetParts *incomingDataStringParts) error {
	packetNumber, err := parsePacketNumber(packetParts.PacketNumber())
	if err != nil {
		return errors.New("parsing packet number part returner err")
	}
	if cast.ToUint(packetParts.PacketNumber()) != packetNumber {
		return errors.New("packet number was not parsed correctly")
	}
	return nil
}

func (s *packetTestSuite) testParsePacketDeviceId(packetParts *incomingDataStringParts) error {
	deviceID, err := parsePacketDeviceID(packetParts.DeviceID())
	if err != nil {
		return errors.New("parsing deviceID was part returner err")
	}
	if cast.ToUint(packetParts.DeviceID()) != deviceID {
		return errors.New("packet deviceID was not parsed correctly")
	}
	return nil
}

func (s *packetTestSuite) getCorrectIncomingDataParts() *incomingDataStringParts {
	values := []string{
		"1.0", "2.00", "3.000", "4", "5", "6", "7", "8", "9", "10",
		"1", "2", "3", "4.4", "5", "6.666666"}
	return newIncomingDataStringParts(
		packetConfig.Token(),
		values,
		"555555555",
		"13",
		"1",
	)
}

func (s *packetTestSuite) composeIncomingBinaryData(
	dataParts *incomingDataStringParts, delimiter, terminator string) []byte {

	packet := dataParts.Token() + delimiter +
		strings.Join(dataParts.Values(), delimiter) + delimiter +
		dataParts.Time() + delimiter +
		dataParts.PacketNumber() + delimiter +
		dataParts.DeviceID() + terminator
	return []byte(packet)
}

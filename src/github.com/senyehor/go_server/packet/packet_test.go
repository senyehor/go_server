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
	correctIncomingData  *incomingDataStringParts
	expectedParsedValues []float64
}

func (s *packetTestSuite) SetupTest() {
	s.expectedParsedValues = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4.4, 5, 6.66666}
}

func (s *packetTestSuite) TestParsingIncomingDataWithCorrectValues() {
	correctData := s.getCorrectIncomingDataParts()
	correctDelimiter := packetConfig.DataDelimiter()
	correctTerminator := packetConfig.DataTerminator()

	parts, err := parseBinaryDataToStringParts(
		s.composeIncomingBinaryData(correctData, string(correctDelimiter), string(correctTerminator)))
	s.NoError(err, "parsing incoming binary data into string parts returned err")
	s.True(correctData.IsEqual(parts), "parsing binary data into string parts went wrong")

	s.NoError(s.testParsePacketValues(parts))
	s.NoError(s.testParsePacketTime(parts))
	s.NoError(s.testParsePacketNumber(parts))
	s.NoError(s.testParsePacketDeviceId(parts))
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

func (s *packetTestSuite) testParsePacketTime(packetParts *incomingDataStringParts) error {
	time, err := parsePacketTime(packetParts.Time())
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

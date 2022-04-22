package data_models

import (
	"fmt"
	"github.com/senyehor/go_server/utils"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"testing"
	"time"
)

func TestPacket(t *testing.T) {
	suite.Run(t, new(appModelsTestSuite))
}

type appModelsTestSuite struct {
	suite.Suite
	correctValues []float64
	correctTimeInterval,
	correctPacketNumber,
	correctDeviceID int
}

func (s *appModelsTestSuite) SetupTest() {
	rand.Seed(time.Now().UnixNano())
	// filling with random negative and positive numbers and ensuring values has correct length
	for i := 0; i < utils.PacketConfig.ValuesCount(); i++ {
		s.correctValues = append(s.correctValues, (rand.Float64()-float64(0.5))*(1_000_000))
	}
	getPositiveInt := func() int {
		result := rand.Int()
		for result = rand.Int(); result < 0; {
			result = rand.Int()
		}
		return result
	}
	s.correctTimeInterval = getPositiveInt()
	s.correctPacketNumber = getPositiveInt()
	s.correctDeviceID = getPositiveInt()
}

func (s *appModelsTestSuite) TestCreatingPacketFromKnowinglyCorrectData() {
	fmt.Print("aaa")
}

//func (s *appModelsTestSuite) TestParsingIncomingDataWithCorrectValues() {
//	correctBinaryData := s.composeIncomingBinaryData(
//		s.correctIncomingDataParts,
//		string(s.correctDelimiter),
//		string(s.correctTerminator),
//	)
//	parts, err := s.testParsingBinaryDataToStringParts(correctBinaryData)
//	s.NoError(err, "parsing binary data to string parts went wrong")
//
//	s.NoError(s.testParsePacketValues(parts), "parsing Packet packetValues went wrong")
//	s.NoError(s.testParsePacketTimeInterval(parts), "parsing Packet time interval went wrong")
//	s.NoError(s.testParsePacketNumber(parts), "parsing Packet number went wrong")
//	s.NoError(s.testParsePacketDeviceId(parts), "parsing Packet device id went wrong")
//}
//
//func (s *appModelsTestSuite) TestParsingToStringPartsWithWrongDelimiter() {
//	binaryDataWithWrongDelimiter := s.composeIncomingBinaryData(
//		s.correctIncomingDataParts,
//		"wrong delimiter",
//		string(s.correctTerminator),
//	)
//	_, err := s.testParsingBinaryDataToStringParts(binaryDataWithWrongDelimiter)
//	s.Error(err, "parsing binary data to string parts with wrong delimiter did not return error")
//}
//
//func (s *appModelsTestSuite) TestParsingToStringPartsWithWrongTerminator() {
//	binaryDataWithWrongTerminator := s.composeIncomingBinaryData(
//		s.correctIncomingDataParts,
//		string(s.correctDelimiter),
//		"wrong terminator",
//	)
//	_, err := s.testParsingBinaryDataToStringParts(binaryDataWithWrongTerminator)
//	s.Error(err, "parsing binary data to string parts with wrong  did not return error")
//}
//
//func (s *appModelsTestSuite) testParsingBinaryDataToStringParts(incomingData []byte) (*incomingDataStringParts, error) {
//	parts, err := parseBinaryDataToStringParts(incomingData)
//	if err != nil {
//		return nil, errors.New("parsing incoming binary data into string parts returned err")
//	}
//	if !s.correctIncomingDataParts.IsEqual(parts) {
//		return nil, errors.New("parsing binary data into string parts went wrong")
//	}
//	return parts, nil
//}
//
//func (s *appModelsTestSuite) testParsePacketValues(packetParts *incomingDataStringParts) error {
//	values, err := parsePacketValues(packetParts.Values())
//	if err != nil {
//		return err
//	}
//	for iterItem := range values.Iterate() {
//		if !utils.CompareFloatsPrecise(iterItem.Value(), s.expectedParsedValues[iterItem.ValuePosition()]) {
//			return errors.New("parsed value does not match expected")
//		}
//	}
//	return nil
//}
//
//func (s *appModelsTestSuite) testParsePacketTimeInterval(packetParts *incomingDataStringParts) error {
//	time, err := parsePacketTimeInterval(packetParts.Time())
//	if err != nil {
//		return errors.New("parsing Packet timeInterval part returner err")
//	}
//	if cast.ToUint(packetParts.Time()) != time {
//		return errors.New("Packet time was not parsed correctly")
//	}
//	return nil
//}
//
//func (s *appModelsTestSuite) testParsePacketNumber(packetParts *incomingDataStringParts) error {
//	packetNumber, err := parsePacketNumber(packetParts.PacketNumber())
//	if err != nil {
//		return errors.New("parsing Packet number part returner err")
//	}
//	if cast.ToUint(packetParts.PacketNumber()) != packetNumber {
//		return errors.New("Packet number was not parsed correctly")
//	}
//	return nil
//}
//
//func (s *appModelsTestSuite) testParsePacketDeviceId(packetParts *incomingDataStringParts) error {
//	deviceID, err := parsePacketDeviceID(packetParts.DeviceID())
//	if err != nil {
//		return errors.New("parsing deviceID was part returner err")
//	}
//	if cast.ToUint(packetParts.DeviceID()) != deviceID {
//		return errors.New("Packet deviceID was not parsed correctly")
//	}
//	return nil
//}
//
//func (s *appModelsTestSuite) getCorrectIncomingDataParts() *incomingDataStringParts {
//	values := []string{
//		"1.0", "2.00", "3.000", "4", "5", "6", "7", "8", "9", "10",
//		"1", "2", "3", "4.4", "5", "6.666666"}
//	return newIncomingDataStringParts(
//		packetConfig.Token(),
//		values,
//		"555555555",
//		"13",
//		"1",
//	)
//}
//
//func (s *appModelsTestSuite) composeIncomingBinaryData(
//	dataParts *incomingDataStringParts, delimiter, terminator string) []byte {
//
//	Packet := dataParts.Token() + delimiter +
//		strings.Join(dataParts.Values(), delimiter) + delimiter +
//		dataParts.Time() + delimiter +
//		dataParts.PacketNumber() + delimiter +
//		dataParts.DeviceID() + terminator
//	return []byte(Packet)
//}

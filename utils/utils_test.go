package utils

import (
	"github.com/stretchr/testify/suite"
	"os"
	"strconv"
	"testing"
)

func TestUtils(t *testing.T) {
	suite.Run(t, new(utilsTestSuite))
}

type utilsTestSuite struct {
	suite.Suite
}

func (u *utilsTestSuite) TestGetStringFromEnv() {
	os.Clearenv()
	key := "KEY"
	value := "VALUE"
	_ = os.Setenv(key, value)
	u.Equal(value, getStringFromEnv(key), "getStringFromEnv returned wrong value")

	u.Panics(
		func() { getStringFromEnv("not_existing_key") },
		"getStringFromEnv did not panic with key that is not in the env",
	)
}

func (u *utilsTestSuite) TestGetBoolFromEnv() {
	os.Clearenv()
	key := "BOOLEAN_KEY"
	boolValue := "true"
	_ = os.Setenv(key, boolValue)
	u.Equal(boolValue, getStringFromEnv(key), "getBoolFromEnv returned wrong value")

	u.Panics(
		func() { getStringFromEnv("wrong_format_bool_value") },
		"getBoolFromEnv did not panic with key that is not in the env",
	)
}

func (u *utilsTestSuite) TestGetRuneFromEnv() {
	runeValue := '^'
	key := "runeValue"
	_ = os.Setenv(key, string(runeValue))

	u.Equal(runeValue, getRuneFromEnv(key), "getRuneFromEnv failed")
	_ = os.Setenv(key, "not rune value")
	u.Panics(func() { getRuneFromEnv(key) }, "getRuneFromEnv does not panic when value is not rune")
}

func (u *utilsTestSuite) TestParseIntConvertToUint() {
	value := 12345
	result, err := ParseIntConvertToUint(strconv.Itoa(value))
	u.NoError(err, "ParseIntConvertToUint returned an error")
	u.Equal(uint(value), result, "ParseIntConvertToUint parsed incorrectly")

	value = -12345
	result, err = ParseIntConvertToUint(strconv.Itoa(value))
	u.Error(err, "ParseIntConvertToUint did not return an error")
	u.Equal(uint(0), result, "ParseIntConvertToUint did not return 0")
}

func (u *utilsTestSuite) TestCompareFloats() {
	u.True(CompareFloatsPrecise(0.0, 0.0))
	u.True(CompareFloatsPrecise(-1111111111.11111, -1111111111.11111))
	u.True(CompareFloatsPrecise(1111111111.11111, 1111111111.11111))

	u.False(CompareFloatsPrecise(0.1, 0.2))
	u.False(CompareFloatsPrecise(-1111111111.11111, 1111111111.11111))
}

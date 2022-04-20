package utils

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

func TestUtils(t *testing.T) {
	suite.Run(t, new(utilsTestSuite))
}

type utilsTestSuite struct {
	suite.Suite
}

func (u *utilsTestSuite) TestGetRuneFromEnv() {
	runeValue := '^'
	key := "runeValue"
	viper.Set(key, string(runeValue))
	u.Equal(runeValue, getRuneFromEnv(key), "getRuneFromEnv failed")

	viper.Set(key, "not rune value")
	u.Panics(func() { getRuneFromEnv(key) }, "getRuneFromEnv does not panic when value is not rune")
}

func (u *utilsTestSuite) TestParseIntConvertToUint() {
	value := 12345
	result, err := ParseIntConvertToUint(strconv.Itoa(value))
	u.NoError(err, "ParseIntConvertToUint returned an error")
	u.Equal(value, result, "ParseIntConvertToUint parsed incorrectly")

	value = -12345
	result, err = ParseIntConvertToUint(strconv.Itoa(value))
	u.Error(err, "ParseIntConvertToUint did not return an error")
	u.Equal(0, result, "ParseIntConvertToUint did not return 0")
}

func (u *utilsTestSuite) TestCompareFloats() {
	u.True(CompareFloats(0.0, 0.0))
	u.True(CompareFloats(-1111111111.11111, -1111111111.11111))
	u.True(CompareFloats(1111111111.11111, 1111111111.11111))

	u.False(CompareFloats(0.1, 0.2))
	u.False(CompareFloats(-1111111111.11111, 1111111111.11111))
}

package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math"
	"strconv"
)

func getRuneFromEnv(key string) rune {
	value := viper.GetString(key)
	if len(value) != 1 {
		log.Errorf("getRuneFromEnv received %v", value)
		panic(errors.New("rune value must be 1 symbol long"))
	}
	return rune(value[0])
}

func getUintFromEnv(key string) uint {
	result, err := ParseIntConvertToUint(viper.GetString(key))
	if err != nil {
		panic("failed to get Uint from env")
	}
	return result
}

func ParseIntConvertToUint(toParse string) (uint, error) {
	result, err := strconv.ParseUint(toParse, 10, 32)
	if err != nil {
		return 0, err
	}
	if result < 0 {
		return 0, errors.New("parsed value is below zero")
	}
	return uint(result), nil
}

func CompareFloats(a, b float64) bool {
	a = math.Abs(a)
	b = math.Abs(b)
	return math.Abs(a-b) < 0.0001
}

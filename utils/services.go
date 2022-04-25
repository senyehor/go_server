package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func getStringFromEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		panic("could not get " + key + " from env")
	}
	return strings.Trim(value, "\"")
}

func getBoolFromEnv(key string) bool {
	value := getStringFromEnv(key)
	result, err := strconv.ParseBool(value)
	if err != nil {
		panic("incorrect format of boolean variable " + key)
	}
	return result
}

func getRuneFromEnv(key string) rune {
	value := getStringFromEnv(key)
	if len(value) != 1 {
		log.Errorf("getRuneFromEnv received %v", value)
		panic(errors.New("rune value must be 1 symbol long"))
	}
	return rune(value[0])
}

func getUintFromEnv(key string) uint {
	result, err := ParseIntConvertToUint(getStringFromEnv(key))
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

func CompareFloatsPrecise(a, b float64) bool {
	if b > a {
		a, b = b, a
	}
	return math.Abs(a-b) < 0.0_000_000_1
}

func RandPositiveInt() int {
	time.Sleep(time.Nanosecond * 100)
	rand.Seed(time.Now().UnixNano())
	result := rand.Int31()
	return int(result)
}

func RandFloat64() float64 {
	time.Sleep(time.Nanosecond * 100)
	rand.Seed(time.Now().UnixNano())
	return (rand.Float64() - float64(0.5)) * (1_000_000.0)
}

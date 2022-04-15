package utils

import (
	"errors"
	"github.com/spf13/viper"
)

func getRuneFromEnv(key string) rune {
	value := viper.GetString(key)
	if len(value) != 1 {
		panic(errors.New("rune value must be 1 symbol long"))
	}
	return rune(value[0])
}

func getUintFromEnv(key string) uint {
	value := viper.GetInt(key)
	if value < 0 {
		panic("value below zero")
	}
	return uint(value)
}

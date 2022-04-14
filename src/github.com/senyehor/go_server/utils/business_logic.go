package utils

import (
	"errors"
	"github.com/spf13/viper"
)

func getRuneFromEnv(key string) rune {
	value := viper.GetString(key)
	if len(value) != 1 {
		panic(errors.New("packet delimiter must be 1 symbol long"))
	}
	return rune(value[0])
}

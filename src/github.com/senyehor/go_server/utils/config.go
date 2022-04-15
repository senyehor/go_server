package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func init() {
	setUpEnv()
}

func GetAppConfig() *appConfig {
	return &appConfig{
		port:  viper.GetString("APP.PORT"),
		debug: viper.GetBool("APP.DEBUG"),
	}
}

func GetPacketConfig() *packetConfig {
	return &packetConfig{
		dataDelimiter:       getRuneFromEnv("PACKET.DELIMITER"),
		dataTerminator:      getRuneFromEnv("PACKET.TERMINATOR"),
		response:            viper.GetString("PACKET.RESPONSE"),
		token:               viper.GetString("PACKET.TOKEN"),
		valuesCount:         uint8(getUintFromEnv("PACKET.VALUES_COUNT")),
		nonValuesPartsCount: uint8(getUintFromEnv("PACKET.NON_VALUES_PARTS_COUNT")),
	}
}

func GetDBConfig() *dbConfig {
	return &dbConfig{
		DBUsername: viper.GetString("DB.USERNAME"),
		DBPassword: viper.GetString("DB.PASSWORD"),
		DBHost:     viper.GetString("DB.HOST"),
		DBPort:     viper.GetString("DB.PORT"),
		DBName:     viper.GetString("DB.NAME"),
	}
}

func setUpEnv() {
	path, _ := os.Getwd()
	// checks what environment app is running in
	viper.SetConfigName("dev_config")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err == nil {
		return
	}
	viper.SetConfigName("prod_config")
	viper.AddConfigPath("/bin/")
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("Failed to find both prod and dev configs")
		panic(err)
	}
	// packet config is the same for all environments
	viper.SetConfigName("packet")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
}

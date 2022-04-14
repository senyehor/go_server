package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func GetAppConfig() *appConfig {
	return &appConfig{
		port:  viper.GetString("APP.PORT"),
		debug: viper.GetBool("APP.DEBUG"),
	}
}

func GetPacketConfig() *packetConfig {
	packetDataDelimiterFromConfig := viper.GetString("PACKET.DELIMITER")
	if len(packetDataDelimiterFromConfig) != 1 {
		panic(errors.New("packet delimiter must be 1 symbol long"))
	}
	packetDataDelimiter := rune(packetDataDelimiterFromConfig[0])
	packetDataTerminatorFromConfig := viper.GetString("PACKET.TERMINATOR")
	if len(packetDataTerminatorFromConfig) != 1 {
		panic(errors.New("packet terminator must be 1 symbol long"))
	}
	packetDataTerminator := rune(packetDataTerminatorFromConfig[0])
	return &packetConfig{
		packetMaxLength:      viper.GetInt("PACKET.MAX_LENGTH"),
		packetDataDelimiter:  packetDataDelimiter,
		packetDataTerminator: packetDataTerminator,
		packetResponse:       viper.GetString("PACKET.RESPONSE"),
		packetToken:          viper.GetString("PACKET.TOKEN"),
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

func SetUpEnv() {
	path, _ := os.Getwd()
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
}

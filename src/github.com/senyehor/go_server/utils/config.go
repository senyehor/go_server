package utils

import (
	"errors"
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
		dataDelimiter:  getRuneFromEnv("PACKET.DELIMITER"),
		dataTerminator: getRuneFromEnv("PACKET.TERMINATOR"),
		response:       viper.GetString("PACKET.RESPONSE"),
		token:          viper.GetString("PACKET.TOKEN"),
		// getUintFromEnv serves as non-negative check
		valuesCount:     int(getUintFromEnv("PACKET.VALUES_COUNT")),
		otherPartsCount: int(getUintFromEnv("PACKET.NON_VALUES_PARTS_COUNT")),
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
	path, found := os.LookupEnv("GO_APP_CONFIG_PATH")
	if !found {
		panic(errors.New("config path environmental variable not found"))
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("packet_config")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Failed to find packet config")
		panic(err)
	}

	// checks what environment app is running in
	viper.SetConfigName("dev_config")
	err = viper.MergeInConfig()
	if err == nil {
		return
	}
	viper.SetConfigName("prod_config")
	err = viper.MergeInConfig()
	if err != nil {
		log.Error("Failed to find both prod and dev configs")
		panic(err)
	}
}

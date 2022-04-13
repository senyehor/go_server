package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type dbConfig struct {
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

// todo check types packetConf

type packetConfig struct {
	packetMaxLength      int    `mapstructure:"PACKET_MAX_LENGTH"`
	packetDataDelimiter  string `mapstructure:"PACKET_DATA_DELIMITER"`
	packetDataTerminator rune   `mapstructure:"PACKET_DATA_TERMINATOR"`
	packetResponse       string `mapstructure:"PACKET_RESPONSE"`
	packetToken          string `mapstructure:"PACKET_TOKEN"`
}

type appConfig struct {
	port  string `mapstructure:"PORT"`
	debug bool   `mapstructure:"DEBUG"`
}

func GetAppConfig() *appConfig {
	setUpEnv()
	return &appConfig{
		port: viper.GetString("APP.PORT"),
	}
}

func GetPacketConfig() *packetConfig {
	setUpEnv()
	return &packetConfig{
		packetMaxLength:      viper.GetInt("PACKET.MAX_LENGTH"),
		packetDataDelimiter:  viper.GetString("PACKET.DELIMITER"),
		packetDataTerminator: rune(viper.GetUint32("PACKET.TERMINATOR")),
		packetResponse:       viper.GetString("PACKET.RESPONSE"),
		packetToken:          viper.GetString("PACKET.TOKEN"),
	}
}

func GetDBConfig() *dbConfig {
	setUpEnv()
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

func (packetConfig *packetConfig) MaxLength() int {
	return packetConfig.packetMaxLength
}

func (packetConfig *packetConfig) DataDelimiter() string {
	return packetConfig.packetDataDelimiter
}
func (packetConfig *packetConfig) DataTerminator() rune {
	return packetConfig.packetDataTerminator
}
func (packetConfig *packetConfig) Response() string {
	return packetConfig.packetResponse
}
func (packetConfig *packetConfig) Token() string {
	return packetConfig.packetToken
}

func (server *appConfig) Port() string {
	return server.port
}

func (DBConfig *dbConfig) Username() string {
	return DBConfig.DBUsername
}
func (DBConfig *dbConfig) Password() string {
	return DBConfig.DBPassword
}
func (DBConfig *dbConfig) Host() string {
	return DBConfig.DBHost
}
func (DBConfig *dbConfig) Port() string {
	return DBConfig.DBPort
}
func (DBConfig *dbConfig) Name() string {
	return DBConfig.DBName
}

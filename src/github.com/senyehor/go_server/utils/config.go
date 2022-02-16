package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type DBConfig struct {
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

// todo check types packetConf

type PacketConfig struct {
	packetMaxLength      int    `mapstructure:"PACKET_MAX_LENGTH"`
	packetDataDelimiter  string `mapstructure:"PACKET_DATA_DELIMITER"`
	packetDataTerminator rune   `mapstructure:"PACKET_DATA_TERMINATOR"`
	packetResponse       string `mapstructure:"PACKET_RESPONSE"`
	packetToken          string `mapstructure:"PACKET_TOKEN"`
}

type ServerConfig struct {
	port string `mapstructure:"PORT"`
}

func GetServerConfig() *ServerConfig {
	setUpEnv()
	return &ServerConfig{
		port: viper.GetString("SERVER.PORT"),
	}
}

func GetPacketConfig() *PacketConfig {
	setUpEnv()
	return &PacketConfig{
		packetMaxLength:      viper.GetInt("PACKET.MAX_LENGTH"),
		packetDataDelimiter:  viper.GetString("PACKET.DELIMITER"),
		packetDataTerminator: rune(viper.GetUint32("PACKET.TERMINATOR")),
		packetResponse:       viper.GetString("PACKET.RESPONSE"),
		packetToken:          viper.GetString("PACKET.TOKEN"),
	}
}

func GetDBConfig() *DBConfig {
	setUpEnv()
	return &DBConfig{
		DBUsername: viper.GetString("DB.USERNAME"),
		DBPassword: viper.GetString("DB.PASSWORD"),
		DBHost:     viper.GetString("DB.HOST"),
		DBPort:     viper.GetString("DB.PORT"),
		DBName:     viper.GetString("DB.NAME"),
	}
}

func setUpEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	path, _ := os.Getwd()
	log.Info("current path is " + path)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("config file was not found")
		panic(err)
	}
}

func (packetConfig *PacketConfig) MaxLength() int {
	return packetConfig.packetMaxLength
}

func (packetConfig *PacketConfig) DataDelimiter() string {
	return packetConfig.packetDataDelimiter
}
func (packetConfig *PacketConfig) DataTerminator() rune {
	return packetConfig.packetDataTerminator
}
func (packetConfig *PacketConfig) Response() string {
	return packetConfig.packetResponse
}
func (packetConfig *PacketConfig) Token() string {
	return packetConfig.packetToken
}

func (server *ServerConfig) Port() string {
	return server.port
}

func (DBConfig *DBConfig) Username() string {
	return DBConfig.DBUsername
}
func (DBConfig *DBConfig) Password() string {
	return DBConfig.DBPassword
}
func (DBConfig *DBConfig) Host() string {
	return DBConfig.DBHost
}
func (DBConfig *DBConfig) Port() string {
	return DBConfig.DBPort
}
func (DBConfig *DBConfig) Name() string {
	return DBConfig.DBName
}

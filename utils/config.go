package utils

func init() {
	PacketConfig = getPacketConfig()
	AppConfig = getAppConfig()
	DBConfig = getDBConfig()
}

var (
	PacketConfig *packetConfig
	AppConfig    *appConfig
	DBConfig     *dbConfig
)

func getAppConfig() *appConfig {
	return &appConfig{
		port:  getStringFromEnv("APP_PORT"),
		debug: getBoolFromEnv("APP_DEBUG"),
	}
}

func getPacketConfig() *packetConfig {
	return &packetConfig{
		dataDelimiter:  getRuneFromEnv("PACKET_DELIMITER"),
		dataTerminator: getRuneFromEnv("PACKET_TERMINATOR"),
		response:       getStringFromEnv("PACKET_RESPONSE"),
		token:          getStringFromEnv("PACKET_TOKEN"),
		// getUintFromEnv serves as non-negative check
		valuesCount:     int(getUintFromEnv("PACKET_VALUES_COUNT")),
		otherPartsCount: int(getUintFromEnv("PACKET_NON_VALUES_PARTS_COUNT")),
	}
}

func getDBConfig() *dbConfig {
	return &dbConfig{
		DBUsername: getStringFromEnv("DB_USERNAME"),
		DBPassword: getStringFromEnv("DB_PASSWORD"),
		DBHost:     getStringFromEnv("DB_HOST"),
		DBPort:     getStringFromEnv("DB_PORT"),
		DBName:     getStringFromEnv("DB_NAME"),
	}
}

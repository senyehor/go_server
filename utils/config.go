package utils

func init() {
	PacketConfig = getPacketConfig()
	AppConfig = getAppConfig()
	DBConfig = getDBConfig()
	ServerControllingConfig = getServerControllingConfig()
	RedisConfig = getRedisConfig()
}

var (
	PacketConfig            *packetConfig
	AppConfig               *appConfig
	DBConfig                *dbConfig
	ServerControllingConfig *serverControllingConfig
	RedisConfig             *redisConfig
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

func getServerControllingConfig() *serverControllingConfig {
	return &serverControllingConfig{
		currentStatusKey:       getStringFromEnv("CURRENT_STATUS_KEY"),
		channelName:            getStringFromEnv("CHANNEL_NAME"),
		resumeListeningCommand: getStringFromEnv("RESUME_LISTENING_COMMAND"),
		stopListeningCommand:   getStringFromEnv("STOP_LISTENING_COMMAND"),
	}
}

func getRedisConfig() *redisConfig {
	return &redisConfig{
		Address:  getStringFromEnv("REDIS_ADDRESS"),
		Password: getStringFromEnv("REDIS_PASSWORD"),
		DB:       int(getUintFromEnv("REDIS_DB")),
	}
}

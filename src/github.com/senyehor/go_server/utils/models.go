package utils

type dbConfig struct {
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

type packetConfig struct {
	packetMaxLength      int    `mapstructure:"PACKET_MAX_LENGTH"`
	packetDataDelimiter  rune   `mapstructure:"PACKET_DATA_DELIMITER"`
	packetDataTerminator rune   `mapstructure:"PACKET_DATA_TERMINATOR"`
	packetResponse       string `mapstructure:"PACKET_RESPONSE"`
	packetToken          string `mapstructure:"PACKET_TOKEN"`
}

type appConfig struct {
	port  string `mapstructure:"PORT"`
	debug bool   `mapstructure:"DEBUG"`
}

func (packetConfig *packetConfig) MaxLength() int {
	return packetConfig.packetMaxLength
}
func (packetConfig *packetConfig) DataDelimiter() rune {
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

func (server *appConfig) ListenAddress() string {
	return "0.0.0.0:" + server.port
}
func (server *appConfig) Debug() bool {
	return server.debug
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

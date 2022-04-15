package utils

type dbConfig struct {
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

type packetConfig struct {
	dataDelimiter       rune   `mapstructure:"PACKET_DATA_DELIMITER"`
	dataTerminator      rune   `mapstructure:"PACKET_DATA_TERMINATOR"`
	response            string `mapstructure:"PACKET_RESPONSE"`
	token               string `mapstructure:"PACKET_TOKEN"`
	valuesCount         uint8  `mapstructure:"PACKET_VALUES_COUNT"`
	nonValuesPartsCount uint8  `mapstructure:"PACKET_NON_VALUES_PARTS_COUNT"`
}

type appConfig struct {
	port  string `mapstructure:"PORT"`
	debug bool   `mapstructure:"DEBUG"`
}

func (packetConfig *packetConfig) DataDelimiter() rune {
	return packetConfig.dataDelimiter
}
func (packetConfig *packetConfig) DataTerminator() rune {
	return packetConfig.dataTerminator
}
func (packetConfig *packetConfig) Response() string {
	return packetConfig.response
}
func (packetConfig *packetConfig) Token() string {
	return packetConfig.token
}
func (packetConfig packetConfig) ValuesCount() uint8 {
	return packetConfig.valuesCount
}

func (packetConfig packetConfig) NonValuesPartsCount() uint8 {
	return packetConfig.nonValuesPartsCount
}

func (app *appConfig) ListenAddress() string {
	return "0.0.0.0:" + app.port
}
func (app *appConfig) Debug() bool {
	return app.debug
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

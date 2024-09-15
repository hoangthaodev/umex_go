package setting

type Config struct {
	Mysql  MysqlSettings `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Server ServerConfig  `mapstructure:"server"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MysqlSettings struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Dbname             string `mapstructure:"database"`
	MaxIdleConns       int    `mapstructure:"maxIdleConns"`
	MaxOpenConns       int    `mapstructure:"maxOpenConns"`
	ConnectMaxLifeTime int    `mapstructure:"connectMaxLifeTime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_age       int    `mapstructure:"max_age"`
	Max_backups   int    `mapstructure:"max_backups"`
	Compress      bool   `mapstructure:"compress"`
}

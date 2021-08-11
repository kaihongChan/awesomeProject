package conf

// GormConf gorm配置
type GormConf struct {
	Driver string `mapstructure:"driver" json:"driver" yaml:"driver"`
}

// MysqlConf mysql配置
type MysqlConf struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	Dbname      string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Username    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	Params      string `mapstructure:"params" json:"params" yaml:"params"`
	LogMode     string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	MaxIdleConn int    `mapstructure:"max-idle-conn" json:"max-idle-conn" yaml:"max-idle-conn"`
	MaxOpenConn int    `mapstructure:"max-open-conn" json:"max-open-conn" yaml:"max-open-conn"`
}

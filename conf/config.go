package conf

type Config struct {
	ServerConf ServerConf `mapstructure:"server" json:"server" yaml:"server"`
	RedisConf  RedisConf  `mapstructure:"redis" json:"redis" yaml:"redis"`
	GormConf   GormConf   `mapstructure:"gorm" json:"gorm" yaml:"gorm"`
	MysqlConf  MysqlConf  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//CasbinConf Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}

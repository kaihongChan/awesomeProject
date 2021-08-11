package conf

type ServerConf struct {
	Port int `mapstructure:"port" json:"port" yaml:"port"`
}

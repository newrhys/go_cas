package config

type Redis struct {
	Addr		string 	`mapstructure:"addr" json:"addr" yaml:"addr"`
	DB       	int    `mapstructure:"db" json:"db" yaml:"db"`
	Password 	string `mapstructure:"password" json:"password" yaml:"password"`
}
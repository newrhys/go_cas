package config

type Server struct {
	Zap 		Zap 	`mapstructure:"zap" json:"zap" yaml:"zap"`
	System 		System 	`mapstructure:"system" json:"system" yaml:"system"`
	JWT 		JWT 	`mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis 		Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql 		Mysql 	`mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local      	Local	`mapstructure:"local" json:"local" yaml:"local"`
	Captcha 	Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Casbin  	Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
}

package config

type Mysql struct {
	Addr			string 	`mapstructure:"addr" json:"addr" yaml:"addr"`
	Dbname      	string 	`mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username    	string 	`mapstructure:"username" json:"username" yaml:"username"`
	Password    	string 	`mapstructure:"password" json:"password" yaml:"password"`
	Charset			string 	`mapstructure:"charset" json:"charset" yaml:"charset"`
	Loc				string 	`mapstructure:"loc" json:"loc" yaml:"loc"`
	MaxIdleConns 	int    	`mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns 	int    	`mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      	bool   	`mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap      	string	`mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}
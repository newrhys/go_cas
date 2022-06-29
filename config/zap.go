package config
type Zap struct {
	Level    	string `mapstructure:"level" json:"host" yaml:"level"`
	Director 	string `mapstructure:"director" json:"director" yaml:"director"`
	MaxSize		int `mapstructure:"max-size" json:"maxSize" yaml:"max-size"`
	MaxBackups	int `mapstructure:"max-backups" json:"maxBackups" yaml:"max-backups"`
	MaxAge		int `mapstructure:"max-age" json:"maxAge" yaml:"max-age"`
	Compress      bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"show-line"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}

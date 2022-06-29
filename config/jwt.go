package config

type JWT struct {
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
}
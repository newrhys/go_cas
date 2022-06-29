package config

type System struct {
	Env				string `mapstructure:"env" json:"env" yaml:"env"`
	ServerPort		string `mapstructure:"server-port" json:"serverPort" yaml:"server-port"`
	DbType			string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       	string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint 	bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	Mode 			string `mapstructure:"mode" json:"mode" yaml:"mode"`
}
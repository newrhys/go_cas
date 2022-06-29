package redis_key

type Jwt struct {
	JwtKey		string 	`mapstructure:"jwt-key" json:"jwtKey" yaml:"jwt-key"`
}
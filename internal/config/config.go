package config

type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Kong struct {
		Host string `yaml:"host"`
		ConsumerRequest string `yaml:"consumer_request"`
		JwtRequest string `yaml:"jwt_request"`
	} `yaml:"kong"`
	Jaeger struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"jaeger"`
}

var Config Configuration
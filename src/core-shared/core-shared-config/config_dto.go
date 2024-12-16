package coresharedconfig

type ConfigDto struct {
	GatewayURL  string `yaml:"gateway_url"`
	ServiceURL  string `yaml:"service_url"`
	ServicePort int    `yaml:"service_port"`

	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`

	LoggingPattern string `yaml:"logging_pattern"`
}

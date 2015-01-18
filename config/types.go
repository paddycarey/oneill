package config

type Configuration struct {
	LogLevel             int                            `yaml:"log_level,omitempty"`
	DefinitionsDirectory string                         `yaml:"definitions_directory,omitempty"`
	DockerApiEndpoint    string                         `yaml:"docker_api_endpoint,omitempty"`
	NginxConfigDirectory string                         `yaml:"nginx_config_directory,omitempty"`
	ServingDomain        string                         `yaml:"serving_domain,omitempty"`
	RegistryCredentials  map[string]RegistryCredentials `yaml:"registry_credentials"`
}

type RegistryCredentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

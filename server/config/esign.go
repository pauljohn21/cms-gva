package config

type Esign struct {
	BaseURL string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
	OpenID  string `mapstructure:"open-id" json:"open-id" yaml:"open-id"`
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

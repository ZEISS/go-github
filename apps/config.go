package apps

import (
	"encoding/json"

	"github.com/go-yaml/yaml"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// NewConfig returns a new Config struct with default values.
func NewConfig() *Config {
	return &Config{}
}

// Config represents the configuration of a GitHub app.
type Config struct {
	// URL is the URL of the GitHub app.
	URL string `yaml:"url" json:"url" envconfig:"GITHUB_APP_URL"`
	// V3APIURL is the URL of the GitHub V3 API. This is the legacy REST API.
	V3APIURL string `yaml:"v3_api_url" json:"v3_api_url" envconfig:"GITHUB_V3_API_URL"`
	// V4APIURL is the URL of the GitHub V4 API. This is the GraphQL API.
	V4APIURL string `yaml:"v4_api_url" json:"v4_api_url" envconfig:"GITHUB_V4_API_URL"`
	// Name is the name of the app.
	Name string `yaml:"name" json:"name" envconfig:"GITHUB_APP_NAME"`
	// IntegrationID is the ID of the app's integration.
	IntegrationID int `yaml:"integration_id" json:"integration_id" envconfig:"GITHUB_APP_INTEGRATION_ID"`
	// WebhookSecret is the secret used to verify webhook requests.
	WebhookSecret string `yaml:"webhook_secret" json:"webhook_secret" envconfig:"GITHUB_APP_WEBHOOK_SECRET"`
	// PrivateKey is the private key used to sign webhook requests.
	PrivateKey string `yaml:"private_key" json:"private_key" envconfig:"GITHUB_APP_PRIVATE_KEY"`
	// ClientID is the ID of the GitHub client.
	ClientID string `yaml:"client_id" json:"client_id" envconfig:"GITHUB_CLIENT_ID"`
	// ClientSecret is the secret of the GitHub client.
	ClientSecret string `yaml:"client_secret" json:"client_secret" envconfig:"GITHUB_CLIENT_SECRET"`
}

// UnmarshalEnv unmarshals environment variables into a Spec struct.
func (c *Config) UnmarshalEnv() error {
	if err := envconfig.Process("", c); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Unmarshal unmarshals JSON data into a Spec struct.
func (c *Config) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, c); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// UnmarshalYAML unmarshals YAML data into a Spec struct.
func (c *Config) UnmarshalYAML(data []byte) error {
	config := Config{}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return errors.WithStack(err)
	}

	*c = config

	return nil
}

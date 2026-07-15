package apps_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeiss/go-github/apps"
)

func TestNewConfig(t *testing.T) {
	config := apps.NewConfig()
	require.NotNil(t, config)
}

func TestUnmarshalEnv(t *testing.T) {
	t.Setenv("GITHUB_APP_URL", "https://github.com/apps/")
	t.Setenv("GITHUB_V3_API_URL", "https://api.github.com/")
	t.Setenv("GITHUB_V4_API_URL", "https://api.github.com/graphql")
	t.Setenv("GITHUB_APP_NAME", "")
	t.Setenv("GITHUB_APP_INTEGRATION_ID", "123456")
	t.Setenv("GITHUB_APP_WEBHOOK_SECRET", "secret")
	t.Setenv("GITHUB_APP_PRIVATE_KEY", "private_key")
	t.Setenv("GITHUB_CLIENT_ID", "client_id")
	t.Setenv("GITHUB_CLIENT_SECRET", "client_secret")

	config := apps.NewConfig()
	require.NotNil(t, config)

	err := config.UnmarshalEnv()
	require.NoError(t, err)

	require.Equal(t, "https://github.com/apps/", config.URL)
	require.Equal(t, "https://api.github.com/", config.V3APIURL)
	require.Equal(t, "https://api.github.com/graphql", config.V4APIURL)
	require.Equal(t, "", config.Name)
	require.Equal(t, 123456, config.IntegrationID)
	require.Equal(t, "secret", config.WebhookSecret)
	require.Equal(t, "private_key", config.PrivateKey)
	require.Equal(t, "client_id", config.ClientID)
	require.Equal(t, "client_secret", config.ClientSecret)
}

func TestUnmarshalYAML(t *testing.T) {
	yamlData := `
url: https://github.com/apps/
v3_api_url: https://api.github.com/
v4_api_url: https://api.github.com/graphql
name: ""
integration_id: 123456
webhook_secret: secret
private_key: private_key
client_id: client_id
client_secret: client_secret
`
	config := apps.NewConfig()
	err := config.UnmarshalYAML([]byte(yamlData))
	require.NoError(t, err)

	require.Equal(t, "https://github.com/apps/", config.URL)
	require.Equal(t, "https://api.github.com/", config.V3APIURL)
	require.Equal(t, "https://api.github.com/graphql", config.V4APIURL)
	require.Equal(t, "", config.Name)
	require.Equal(t, 123456, config.IntegrationID)
	require.Equal(t, "secret", config.WebhookSecret)
	require.Equal(t, "private_key", config.PrivateKey)
	require.Equal(t, "client_id", config.ClientID)
	require.Equal(t, "client_secret", config.ClientSecret)
}

func TestUnmarshal(t *testing.T) {
	config := apps.NewConfig()

	jsonData := `
{
	"url": "https://github.com/apps/",
	"v3_api_url": "https://api.github.com/",
	"v4_api_url": "https://api.github.com/graphql",
	"name": "",
	"integration_id": 123456,
	"webhook_secret": "secret",
	"private_key": "private_key",
	"client_id": "client_id",
	"client_secret": "client_secret"
}`
	err := config.Unmarshal([]byte(jsonData))
	require.NoError(t, err)

	require.Equal(t, "https://github.com/apps/", config.URL)
	require.Equal(t, "https://api.github.com/", config.V3APIURL)
	require.Equal(t, "https://api.github.com/graphql", config.V4APIURL)
	require.Equal(t, "", config.Name)
	require.Equal(t, 123456, config.IntegrationID)
	require.Equal(t, "secret", config.WebhookSecret)
	require.Equal(t, "private_key", config.PrivateKey)
	require.Equal(t, "client_id", config.ClientID)
	require.Equal(t, "client_secret", config.ClientSecret)
}

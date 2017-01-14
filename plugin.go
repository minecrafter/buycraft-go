package buycraft

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PluginClient defines a Buycraft API client.
type PluginClient struct {
	secret string
}

const (
	apiEndpoint  = "https://plugin.buycraft.net"
	apiUserAgent = "buycraft-go (+https://github.com/minecrafter/buycraft-go)"
)

// NewPluginClient creates a new PluginClient.
func NewPluginClient(secret string) *PluginClient {
	return &PluginClient{
		secret: secret,
	}
}

func (pc *PluginClient) pluginGet(endpoint string, information interface{}) error {
	req, err := http.NewRequest(http.MethodGet, apiEndpoint+endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Buycraft-Secret", pc.secret)
	req.Header.Set("User-Agent", apiUserAgent)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("Unexpected content-type %s", res.Header.Get("Content-Type"))
	}

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode == http.StatusOK {
		if err = decoder.Decode(&information); err != nil {
			return err
		}
		return nil
	}

	// Encountered an error
	var pluginErr PluginError
	if err = decoder.Decode(&pluginErr); err != nil {
		return err
	}

	return fmt.Errorf("Buycraft plugin error %d: %s", pluginErr.Code, pluginErr.Message)
}

// Information returns information about the client.
func (pc *PluginClient) Information() (PluginInformation, error) {
	var pi PluginInformation
	err := pc.pluginGet("/information", &pi)
	return pi, err
}

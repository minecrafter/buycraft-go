package buycraft

// PluginError indicates an error returned by the Buycraft plugin API.
type PluginError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

// PluginInformation describes the information for the server.
type PluginInformation struct {
	Account PluginAccount
	Server  PluginServer
}

type PluginAccount struct {
	ID         int
	Domain     string
	Name       string
	Currency   PluginCurrency
	OnlineMode bool `json:"online_mode"`
}

type PluginCurrency struct {
	ISO4217 string `json:"iso_4217"`
	Symbol  string
}

type PluginServer struct {
	ID   int
	Name string
}

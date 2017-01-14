package buycraft

// PluginError indicates an error returned by the Buycraft plugin API.
type PluginError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

// PluginInformation describes the information for the server.
type PluginInformation struct {
	Server PluginServer
}

type PluginServer struct {
	ID   int
	Name string
}

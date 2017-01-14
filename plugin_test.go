package buycraft

import (
	"testing"
)

const (
	// This key points to a completely empty webstore and is not used
	// in production.
	//
	// I trust that nobody decides to do something stupid with this.
	serverSecret = "f2c20664a982c897b28b9dc8b2652439a45c4510"
)

func TestInformation(t *testing.T) {
	client := NewPluginClient(serverSecret)
	information, err := client.Information()
	if err != nil {
		t.Errorf("Unable to get information: %s", err.Error())
	}

	if information.Server.Name != "Buycraftxtestsuite" {
		t.Errorf("Server name is unexpected; wanted 'Buycraftxtestsuite', got '%s'", information.Server.Name)
	}
}

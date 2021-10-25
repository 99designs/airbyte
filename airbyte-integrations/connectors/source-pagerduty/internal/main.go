package internal

import (
	_ "embed"
	"encoding/json"
	"github.com/PagerDuty/go-pagerduty"
)

//go:embed connection_specification.json
var connectorSpecificationBytes []byte

func Specification() (AirbyteMessage, error) {
	var connectorSpecification ConnectorSpecification

	err := json.Unmarshal(connectorSpecificationBytes, &connectorSpecification)

	message := AirbyteMessage{
		MessageType: Spec,
		Spec: &connectorSpecification,
	}

	return message, err
}

func Check(config Config) (AirbyteMessage, error) {
	client := pagerduty.NewClient(config.ApiToken)
	var opts pagerduty.ListAddonOptions

	_, err := client.ListAddons(opts)

	var connectionStatus AirbyteConnectionStatus

	message := AirbyteMessage{
		MessageType: ConnectionStatus,
		ConnectionStatus: &connectionStatus,
	}

	if err != nil {
		connectionStatus = AirbyteConnectionStatus{
			Status: Failed,
			Message: err.Error(),
		}
	} else {
		connectionStatus = AirbyteConnectionStatus{
			Status: Succeeded,
			Message: "Yay!",
		}
	}

	return message, nil
}

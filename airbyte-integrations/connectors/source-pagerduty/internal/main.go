package internal

import (
	_ "embed"
	"encoding/json"
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
	connectionStatus := AirbyteConnectionStatus{
		Status: Succeeded,
		Message: "Yay!",
	}

	message := AirbyteMessage{
		MessageType: ConnectionStatus,
		ConnectionStatus: &connectionStatus,
	}

	return message, nil
}

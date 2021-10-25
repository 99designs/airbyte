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
	message := AirbyteMessage{
		MessageType: ConnectionStatus,
	}
	return message, nil
}

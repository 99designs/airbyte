package internal

import (
	_ "embed"
	"encoding/json"
)

//go:embed connection_specification.json
var connectorSpecificationBytes []byte

type ConnectorSpecification struct {
	DocumentationUrl string
	ConnectionSpecification interface{}
}

func Specification() (ConnectorSpecification, error) {
	var connectorSpecification ConnectorSpecification

	err	:= json.Unmarshal(connectorSpecificationBytes, &connectorSpecification)

	return connectorSpecification, err
}

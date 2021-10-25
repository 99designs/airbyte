package internal

import (
	_ "embed"
	"encoding/json"
	"errors"
	"time"
)

//go:embed connection_specification.json
var connectorSpecificationBytes []byte

type ConnectorSpecification struct {
	DocumentationUrl        string
	ConnectionSpecification interface{}
}

func Specification() (ConnectorSpecification, error) {
	var connectorSpecification ConnectorSpecification

	err := json.Unmarshal(connectorSpecificationBytes, &connectorSpecification)

	return connectorSpecification, err
}

type Config struct {
	apiToken string `json:"api_token"`
	// TODO: We can unpack this as a date
	startDate time.Time `json:"start_date"`
}

type Status int

const (
	Failed Status = iota
	Succeeded
)

func (status *Status) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case "succeeded":
		*status = Succeeded
	case "failed":
		*status = Failed
	default:
		errors.New("Unknown status")
	}

	return nil
}

func (status Status) MarshalJSON() ([]byte, error) {
	var s string
	switch status {
	case Succeeded:
		s = "succeeded"
	case Failed:
		s = "failed"
	}

	return json.Marshal(s)
}

type AirbyteConnectionStatus struct {
	// TODO: We can unpack this as an enum
	status  Status
	message string
}

func Check(config Config) (AirbyteConnectionStatus, error) {
	var status AirbyteConnectionStatus
	return status, nil
}

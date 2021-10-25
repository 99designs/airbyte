package internal

import (
	"encoding/json"
	"time"
	"errors"
)

type ConnectorSpecification struct {
	DocumentationUrl string `json:"documentationUrl"`
	ConnectionSpecification interface{} `json:"connectionSpecification"`
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

type MessageType int

const (
	Record MessageType = iota
	State
	Log
	Spec
	ConnectionStatus
	Catalog
)

func (m *MessageType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "RECORD":
		*m = Record
	case "STATE":
		*m = State
	case "LOG":
		*m = Log
	case "SPEC":
		*m = Spec
	case "CONNECTION_STATUS":
		*m = ConnectionStatus
	case "CATALOG":
		*m = Catalog
	}
	return nil
}

func (m MessageType) MarshalJSON() ([]byte, error) {
	var s string
	switch m {
	case Record:
		s = "RECORD"
	case State:
		s = "STATE"
	case Log:
		s = "LOG"
	case Spec:
		s = "SPEC"
	case ConnectionStatus:
		s = "CONNECTION_STATUS"
	case Catalog:
		s = "CATALOG"
	}

	return json.Marshal(s)
}

type AirbyteMessage struct {
	MessageType MessageType `json:"type"`
	// log AirbyteLogMessage
	Spec *ConnectorSpecification `json:"spec,omitempty"`
	ConnectionStatus *AirbyteConnectionStatus `json:"connectionStatus,omitempty"`
	// catalog AirbyteCatalog
	// record AirbyteRecordMessage
	// state AirbyteStateMessae
}

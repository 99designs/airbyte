package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/json"
	"time"
)

func TestStatus(t *testing.T) {
	blob := `["SUCCEEDED","FAILED"]`
	list := []Status{Succeeded, Failed}

	var unmarshalResult []Status
	err := json.Unmarshal([]byte(blob), &unmarshalResult)

	assert.Nil(t, err, "Should be a valid json string")
	assert.Equal(t, list, unmarshalResult)

	marshalResult, err := json.Marshal(list)

	assert.Nil(t, err, "Should be able to unmarshal")
	assert.Equal(t, []byte(blob), marshalResult)
}

func TestMessageType(t *testing.T) {
	blob := `["RECORD","STATE","LOG","SPEC","CONNECTION_STATUS","CATALOG"]`
	list := []MessageType{Record, State, Log, Spec, ConnectionStatus, Catalog}

	var unmarshalResult []MessageType
	err := json.Unmarshal([]byte(blob), &unmarshalResult)

	assert.Nil(t, err, "Should be a valid json string")
	assert.Equal(t, list, unmarshalResult)

	marshalResult, err := json.Marshal(list)

	assert.Nil(t, err, "Should be able to unmarshal")
	assert.Equal(t, blob, string(marshalResult))
}

func TestConfig(t *testing.T) {
	blob := `{"start_date":"2021-01-01T00:00:00Z","api_token":"abc123"}`
	tim, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

	config := Config{
		StartDate: tim,
		ApiToken: "abc123",
	}

	var unmarshalResult Config
	err := json.Unmarshal([]byte(blob), &unmarshalResult)

	assert.Nil(t, err)
	assert.Equal(t, config, unmarshalResult)

	_, err = json.Marshal(config)

	assert.Nil(t, err)
}

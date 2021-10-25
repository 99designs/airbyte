package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/json"
)

func TestStatus(t *testing.T) {
	blob := `["succeeded","failed"]`
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
	assert.Equal(t, string([]byte(blob)), string(marshalResult))
}

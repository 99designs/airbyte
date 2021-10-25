package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/json"
)

func TestSpecification(t *testing.T) {
	spec, err := Specification()

	assert.Nil(t, err, "Shouldn't return an error")
	assert.NotZero(t, spec.DocumentationUrl, "Always provide documentation")
	assert.NotZero(
		t,
		spec.ConnectionSpecification,
		"Always provide specificification",
	)
}

func TestStatus(t *testing.T) {
	blob := `["succeeded","failed"]`
	statusList := []Status{Succeeded, Failed}

	var unmarshalResult []Status
	err := json.Unmarshal([]byte(blob), &unmarshalResult)

	assert.Nil(t, err, "Should be a valid json string")
	assert.Equal(t, statusList, unmarshalResult)

	marshalResult, err := json.Marshal([]Status{Succeeded, Failed})

	assert.Nil(t, err, "Should be able to unmarshal")
	assert.Equal(t, []byte(blob), marshalResult)
}

// func TestCheck(t *testing.T) {
// 	config := Config()
// 	status, err := Check(config)

// 	assert.Nil(t, err, "Shouldn't return an error")
// }

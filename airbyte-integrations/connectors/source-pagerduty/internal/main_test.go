package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSpecification(t *testing.T) {
	m, err := Specification()

	assert.Nil(t, err, "Shouldn't return an error")
	assert.Equal(t, Spec, m.messageType)
	assert.NotZero(t, m.spec)
	assert.NotZero(t, m.spec.DocumentationUrl, "Always provide documentation")
	assert.NotZero(
		t,
		m.spec.ConnectionSpecification,
		"Always provide specificification",
	)
}

func TestCheck(t *testing.T) {
	config := Config{"token", time.Now()}
	m, err := Check(config)

	assert.Nil(t, err, "Shouldn't return an error")
	assert.Equal(t, ConnectionStatus, m.messageType)
	assert.NotZero(t, m.connectionStatus)
}

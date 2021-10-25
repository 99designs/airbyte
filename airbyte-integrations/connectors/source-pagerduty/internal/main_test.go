package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSpecification(t *testing.T) {
	m, err := Specification()

	assert.Nil(t, err, "Shouldn't return an error")
	assert.Equal(t, Spec, m.MessageType)
	assert.NotZero(t, m.Spec)
	assert.NotZero(t, m.Spec.DocumentationUrl, "Always provide documentation")
	assert.NotZero(
		t,
		m.Spec.ConnectionSpecification,
		"Always provide specificification",
	)
}

func TestCheck(t *testing.T) {
	config := Config{"token", time.Now()}
	m, err := Check(config)

	assert.Nil(t, err, "Shouldn't return an error")
	assert.Equal(t, ConnectionStatus, m.MessageType)
	assert.NotZero(t, m.ConnectionStatus)
}

func TestDiscover(t *testing.T) {
	config := Config{"token", time.Now()}
	m, err := Discover(config)

	assert.Nil(t, err, "Shouldn't return an error")
	assert.Equal(t, Catalog, m.MessageType)
	assert.NotZero(t, m.Catalog)
}

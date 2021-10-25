package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSpecification(t *testing.T) {
	spec, err := Specification()

	assert.Nil(t, err, "Shouldn't return an error")
	assert.NotZero(t, spec.DocumentationUrl)
	assert.NotZero(t, spec.ConnectionSpecification)
}

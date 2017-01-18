package constants_test

import (
	"testing"

	"github.com/Akagi201/swiggo/constants"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, 42, constants.ICONST)
	assert.Equal(t, 2.1828, constants.FCONST)
	// assert.Equal(t, 'x', constants.CCONST)
	// assert.Equal(t, '\n', constants.CCONST2)
	assert.Equal(t, "Hello World", constants.SCONST)
	assert.Equal(t, 48.5484, constants.EXPR)
	assert.Equal(t, 37, constants.Iconst)
	assert.Equal(t, 3.14, constants.Fconst)
}

package eventstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	assert.NotNil(t, new(baseData))
	assert.NotNil(t, new(baseSchema))
}

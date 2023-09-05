package eventstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	assert.NotNil(t, new(Options))
}

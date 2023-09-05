package eventstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeta(t *testing.T) {
	assert.NotNil(t, new(Meta))
}

package eventstream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := NewClient()

	assert.NotNil(t, client)
	assert.NotNil(t, client.httpClient)
	assert.Equal(t, url, client.url)
	assert.Equal(t, backoffTime, client.backoffTime)
	assert.Equal(t, pageDeleteURL, client.options.PageDeleteURL)
	assert.Equal(t, pageMoveURL, client.options.PageMoveURL)
	assert.Equal(t, revisionCreateURL, client.options.RevisionCreateURL)
	assert.Equal(t, revisionScoreURL, client.options.RevisionScoreURL)
	assert.Equal(t, revisionVisibilityChangeURL, client.options.RevisionVisibilityChangeURL)
}

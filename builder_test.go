package eventstream

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const builderTestURL = "http://localhost:5000"
const builderTestBackoffTime = time.Second * 1
const builderTestPageDeleteURL = "/page-delete"
const builderTestPageMoveURL = "/page-move"
const builderTestRevisionCreateURL = "/revision-create"
const builderTestPageCreateURL = "/page-create"
const builderTestRevisionScoreURL = "/revision-score"
const builderTestRevisionVisibilityChangeURL = "/revision-visibility-change"

func TestBuilder(t *testing.T) {
	options := &Options{
		builderTestPageCreateURL,
		builderTestPageDeleteURL,
		builderTestPageMoveURL,
		builderTestRevisionCreateURL,
		builderTestRevisionScoreURL,
		builderTestRevisionVisibilityChangeURL,
	}
	httpClient := http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
	}

	client := NewBuilder().
		URL(builderTestURL).
		HTTPClient(&httpClient).
		BackoffTime(builderTestBackoffTime).
		Options(options).
		Build()

	assert.NotNil(t, client)
	assert.NotNil(t, client.httpClient)
	assert.NotNil(t, client.httpClient.Transport)
	assert.Equal(t, builderTestBackoffTime, client.backoffTime)
	assert.Equal(t, builderTestURL, client.url)
	assert.Equal(t, builderTestPageDeleteURL, client.options.PageDeleteURL)
	assert.Equal(t, builderTestPageMoveURL, client.options.PageMoveURL)
	assert.Equal(t, builderTestRevisionCreateURL, client.options.RevisionCreateURL)
	assert.Equal(t, builderTestPageCreateURL, client.options.PageCreateURL)
	assert.Equal(t, builderTestRevisionScoreURL, client.options.RevisionScoreURL)
	assert.Equal(t, builderTestRevisionVisibilityChangeURL, client.options.RevisionVisibilityChangeURL)
}

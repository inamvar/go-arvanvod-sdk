package arvanvod_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVideo(t *testing.T) {
	client := getClient()
	videoId := "47b71ff7-416d-4c9d-8d21-02cba7ea4f56"
	resp, err := client.GetVideo(context.Background(), videoId)
	if assert.NoError(t, err) {
		assert.NotNil(t, resp.Data)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Data.Title)
		assert.Equal(t, videoId, resp.Data.Id)
	}
}

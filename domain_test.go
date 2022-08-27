package arvanvod_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDomain(t *testing.T) {

	client := getClient()

	result, err := client.GetDomain(context.Background())

	if assert.NoError(t, err) {
		assert.NotNil(t, result)
	}

}

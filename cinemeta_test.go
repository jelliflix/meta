package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCinemeta(t *testing.T) {
	client := NewCinemeta(DefaultOptions)

	movies := map[string]bool{
		"tt9170516": false,
		"tt91705":   true,
	}
	for id, empty := range movies {
		meta, err := client.GetMovie(id)
		if empty {
			assert.Empty(t, meta)
			assert.Error(t, err)
		} else {
			assert.NotEmpty(t, meta)
			assert.NoError(t, err)
		}
	}

	series := map[string]bool{
		"tt11280740": false,
		"tt11280":    true,
	}
	for id, empty := range series {
		meta, err := client.GetSeries(id)
		if empty {
			assert.Empty(t, meta)
			assert.Error(t, err)
		} else {
			assert.NotEmpty(t, meta)
			assert.NoError(t, err)
		}
	}
}

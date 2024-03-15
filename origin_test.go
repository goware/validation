package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrigin(t *testing.T) {
	t.Run("invalid origin", func(t *testing.T) {
		_, err := NewOrigin("invalid")
		assert.Error(t, err)

		_, err = NewOrigin("http://example.com/invalid")
		assert.ErrorIs(t, err, ErrOriginPath)

		_, err = NewOrigin("http://example.com/?invalid")
		assert.ErrorIs(t, err, ErrOriginQuery)

		_, err = NewOrigin("http://example.com/#invalid")
		assert.ErrorIs(t, err, ErrOriginFragment)

		_, err = NewOrigin("http://example.com/invalid?invalid#invalid")
		assert.ErrorIs(t, err, ErrOriginPath)
		assert.ErrorIs(t, err, ErrOriginQuery)
		assert.ErrorIs(t, err, ErrOriginFragment)
	})
	t.Run("valid origin", func(t *testing.T) {
		origin, err := NewOrigin("http://example.com")
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com", origin.String())

		origin, err = NewOrigin("http://example.com/")
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com", origin.String())

		origin, err = NewOrigin("http://*.example.com/")
		assert.NoError(t, err)
		assert.Equal(t, "http://*.example.com", origin.String())
	})

}

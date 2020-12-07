package constant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLimit(t *testing.T) {
	limit := GetLimit(20)
	assert.Equal(t, 20, limit)

	limit = GetLimit(60)
	assert.Equal(t, 60, limit)

	limit = GetLimit(10)
	assert.Equal(t, 0, limit)
}

func TestCheckLimit(t *testing.T) {
	isLimit := Check(20)
	assert.Equal(t, true, isLimit)

	isLimit = Check(10)
	assert.Equal(t, false, isLimit)
}

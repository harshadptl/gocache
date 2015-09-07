package gocache

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_SetData(t *testing.T) {
	SetData("key", "value", 100)

	val, b := GetData("key")

	assert.Equal(t, true, b)
	assert.Equal(t, "value", val)

	time.Sleep(100 * time.Millisecond)

	val, b = GetData("key")

	assert.Equal(t, false, b)
	assert.Equal(t, "", val)
}

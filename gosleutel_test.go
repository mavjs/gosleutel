package gosleutel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var acc = "github"

func TestCreateSalt(t *testing.T) {
	assert := assert.New(t)

	genSalt := createSalt(acc)
	want := []byte("$sleutel#github#sleutel$")

	assert.Equal(want, genSalt, "Expected equal value")
}

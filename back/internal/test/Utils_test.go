package test

import (
	"testing"

	"github.com/Carlosabdoamaral/go_chatbot_backend/internal/utils"
	"github.com/stretchr/testify/assert"
)

func Test_CapitalizeWord(t *testing.T) {
	a := utils.CapitalizeWord("FOO")
	b := utils.CapitalizeWord("bar")

	assert.Equal(t, "Foo", a)
	assert.Equal(t, "Bar", b)
}

func Test_GetFileCategory(t *testing.T) {
	s := utils.GetFileCategory("/path/to/foo.json")
	assert.Equal(t, "foo", s)
}

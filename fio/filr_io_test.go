package fio

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestNewFileIOManager(t *testing.T) {
	fio, err := NewFileIOManager(filepath.Join("/tmp", "a.data"))
	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

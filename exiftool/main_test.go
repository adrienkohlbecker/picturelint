package exiftool

import (
	"testing"

	"gotest.tools/assert"
)

func TestRead(t *testing.T) {

	metadata, err := Read("testdata/mire.jpg")
	t.Log(metadata)

	assert.NilError(t, err)
	assert.Equal(t, metadata["SourceFile"], "testdata/mire.jpg")

}

package cases

import (
	"testing"

	"github.com/adrienkohlbecker/picturelint/validators"
	"gotest.tools/assert"
)

func TestEXIFDateTimeOriginal(t *testing.T) {

	assert.Equal(t, *CaseResult(t, "testdata/mire.jpg", "has EXIF:DateTimeOriginal set"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has EXIF:DateTimeOriginal set"), validators.StatusSuccess)

}

func TestEXIFModifyDate(t *testing.T) {

	assert.Equal(t, *CaseResult(t, "testdata/mire.jpg", "has EXIF:ModifyDate set"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has EXIF:ModifyDate set"), validators.StatusSuccess)

	assert.Equal(t, *CaseResult(t, "testdata/mire+times_different_from_base.jpg", "has EXIF:ModifyDate set to EXIF:DateTimeOriginal"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has EXIF:ModifyDate set to EXIF:DateTimeOriginal"), validators.StatusSuccess)

}

func TestXMPModifyDate(t *testing.T) {

	assert.Equal(t, *CaseResult(t, "testdata/mire.jpg", "has XMP:ModifyDate set"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has XMP:ModifyDate set"), validators.StatusSuccess)

	assert.Equal(t, *CaseResult(t, "testdata/mire+times_different_from_base.jpg", "has XMP:ModifyDate set to EXIF:DateTimeOriginal"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has XMP:ModifyDate set to EXIF:DateTimeOriginal"), validators.StatusSuccess)

}

func TestXMPDateCreated(t *testing.T) {

	assert.Equal(t, *CaseResult(t, "testdata/mire.jpg", "has XMP:DateCreated set"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has XMP:DateCreated set"), validators.StatusSuccess)

	assert.Equal(t, *CaseResult(t, "testdata/mire+times_different_from_base.jpg", "has XMP:DateCreated set to EXIF:DateTimeOriginal"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has XMP:DateCreated set to EXIF:DateTimeOriginal"), validators.StatusSuccess)

}

func TestIPTCTimeCreated(t *testing.T) {

	assert.Equal(t, *CaseResult(t, "testdata/mire.jpg", "has IPTC:TimeCreated set"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has IPTC:TimeCreated set"), validators.StatusSuccess)

	assert.Equal(t, *CaseResult(t, "testdata/mire+times_different_from_base.jpg", "has IPTC:TimeCreated set to EXIF:DateTimeOriginal"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has IPTC:TimeCreated set to EXIF:DateTimeOriginal"), validators.StatusSuccess)

}

func TestIPTCDateCreated(t *testing.T) {

	assert.Equal(t, *CaseResult(t, "testdata/mire.jpg", "has IPTC:DateCreated set"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has IPTC:DateCreated set"), validators.StatusSuccess)

	assert.Equal(t, *CaseResult(t, "testdata/mire+times_different_from_base.jpg", "has IPTC:DateCreated set to EXIF:DateTimeOriginal"), validators.StatusFailed)
	assert.Equal(t, *CaseResult(t, "testdata/mire+edit_capture_time.jpg", "has IPTC:DateCreated set to EXIF:DateTimeOriginal"), validators.StatusSuccess)

}

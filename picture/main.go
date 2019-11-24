package picture

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adrienkohlbecker/picturelint/exiftool"
)

type Picture struct {
	Path        string
	rawMetadata map[string]interface{}
}

// Includes the dot because that's what filepath.Ext returns
var ImageExtensions = []string{".cr2", ".dng", ".heic", ".jpeg", ".jpg", ".nef", ".png", ".raf", ".tif"}
var VideoExtensions = []string{".mov", ".mp4", ".avi", ".m4v", ".mts"}

func Load(path string) (*Picture, error) {

	p := &Picture{Path: path}
	err := p.ReadMetadata()
	if err != nil {
		return nil, err
	}

	return p, nil

}

func (p *Picture) ReadMetadata() error {

	result, err := exiftool.Read(p.Path)
	if err != nil {
		return err
	}

	p.rawMetadata = result
	return nil

}

func (p *Picture) HasMetadata() bool {

	return len(p.rawMetadata) > 0

}

func (p *Picture) IsVideo() bool {

	return Contains(VideoExtensions, strings.ToLower(filepath.Ext(p.Path)))

}

func (p *Picture) IsVideoPreview() bool {

	if strings.ToLower(filepath.Ext(p.Path)) == ".png" {
		base := strings.TrimSuffix(p.Path, ".png")
		for _, ext := range VideoExtensions {
			if FileExists(base+ext) || FileExists(base+strings.ToUpper(ext)) {
				return true
			}
		}

	}

	return false

}

func (p *Picture) Filename() string {
	return filepath.Base(p.Path)
}

func (p *Picture) Album() string {
	return filepath.Base(filepath.Dir(p.Path))
}

func (p *Picture) YearAlbum() string {
	return filepath.Base(filepath.Dir(filepath.Dir(p.Path)))
}

func (p *Picture) EXIFDateTimeOriginal() string {
	return p.ReadStringField("EXIF:DateTimeOriginal")
}

func (p *Picture) EXIFOffsetTime() string {
	return p.ReadStringField("EXIF:OffsetTime")
}

func (p *Picture) EXIFCreateDate() string {
	return p.ReadStringField("EXIF:CreateDate")
}

func (p *Picture) EXIFModifyDate() string {
	return p.ReadStringField("EXIF:ModifyDate")
}

func (p *Picture) XMPDateTimeOriginal() string {
	return p.ReadStringField("XMP:DateTimeOriginal")
}

func (p *Picture) XMPDateCreated() string {
	return p.ReadStringField("XMP:DateCreated")
}

func (p *Picture) XMPCreateDate() string {
	return p.ReadStringField("XMP:CreateDate")
}

func (p *Picture) XMPModifyDate() string {
	return p.ReadStringField("XMP:ModifyDate")
}

func (p *Picture) XMPXOffsetTime() string {
	return p.ReadStringField("XMP:XOffsetTime")
}

func (p *Picture) IPTCDateCreated() string {
	return p.ReadStringField("IPTC:DateCreated")
}

func (p *Picture) IPTCTimeCreated() string {
	return p.ReadStringField("IPTC:TimeCreated")
}

func (p *Picture) QuickTimeCreationDate() string {
	return p.ReadStringField("QuickTime:CreationDate")
}

func (p *Picture) ParsedXMPDateTimeOriginal() (time.Time, error) {
	return time.Parse("2006:01:02 15:04:05+07:00", p.XMPDateTimeOriginal())
}

func (p *Picture) ReadStringField(field string) string {

	raw, ok := p.rawMetadata[field]
	if !ok {
		return ""
	}

	value, ok := raw.(string)
	if !ok {
		return ""
	}

	return value

}

func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

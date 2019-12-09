package cases

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/adrienkohlbecker/picturelint/picture"
	"github.com/adrienkohlbecker/picturelint/validators"
)

func init() {

	field("EXIF:DateTimeOriginal", func(f string) {

		it(f, "is set", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos use XMP:DateTimeOriginal and don't support EXIF
				c.Skip()
				return
			}

			if p.EXIFDateTimeOriginal() == "" {
				c.Fail()
			}

		})

	})

	field("EXIF:OffsetTime", func(f string) {

		// TODO: Update logic to verify timezone in XMP metadata for pictures

		it(f, "is set (TODO)", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos don't support EXIF
				c.Skip()
				return
			}

			if p.EXIFOffsetTime() == "" {
				c.Fail()
			}

		})

	})

	for _, f := range []string{"EXIF:CreateDate", "EXIF:ModifyDate"} {

		it(f, "is set", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos don't support EXIF
				c.Skip()
				return
			}

			if p.EXIFCreateDate() == "" {
				c.Fail()
			}

		})

		it(f, "is set to EXIF:DateTimeOriginal", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos don't support EXIF
				c.Skip()
				return
			}

			if p.EXIFCreateDate() == "" {
				c.Skip()
				return
			}

			if p.EXIFCreateDate() != p.EXIFDateTimeOriginal() {
				c.Fail()
			}

		})

	}

	field("XMP:DateTimeOriginal", func(f string) {

		it(f, "is set", func(c *validators.Case, p *picture.Picture) {

			if !p.IsVideo() {
				// images use EXIF:DateTimeOriginal as reference
				c.Skip()
				return
			}

			if p.XMPDateTimeOriginal() == "" {
				c.Fail()
			}

		})

		it(f, "has time zone information", func(c *validators.Case, p *picture.Picture) {

			if !p.IsVideo() {
				// images use EXIF:DateTimeOriginal as reference
				c.Skip()
				return
			}

			if p.XMPDateTimeOriginal() == "" {
				c.Skip()
				return
			}

			matched, err := regexp.MatchString(`^\d{4}:\d{2}:\d{2} \d{2}:\d{2}:\d{2}[+|-]\d{2}:\d{2}$`, p.XMPDateTimeOriginal())
			if err != nil || !matched {
				c.Fail()
			}

		})

	})

	for _, f := range []string{"XMP:CreateDate", "XMP:ModifyDate", "XMP:DateCreated"} {

		it(f, "is set", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideoPreview() {
				// we don't manage video previews
				c.Skip()
				return
			}

			if p.XMPModifyDate() == "" {
				c.Fail()
			}

		})

		it(f, "is set to EXIF:DateTimeOriginal", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos use XMP:DateTimeOriginal as reference
				c.Skip()
				return
			}

			if p.XMPModifyDate() == "" {
				c.Skip()
				return
			}

			if p.XMPModifyDate() != p.EXIFDateTimeOriginal() {
				c.Fail()
			}

		})

		it(f, "is set to XMP:DateTimeOriginal", func(c *validators.Case, p *picture.Picture) {

			if !p.IsVideo() {
				// photos use EXIF:DateTimeOriginal as reference
				c.Skip()
				return
			}

			if p.XMPModifyDate() == "" || p.XMPDateTimeOriginal() == "" {
				c.Skip()
				return
			}

			if p.XMPModifyDate() != p.XMPDateTimeOriginal() {
				c.Fail()
			}

		})

	}

	field("IPTC:DateCreated", func(f string) {

		it(f, "is set", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos don't support IPTC
				c.Skip()
				return
			}

			if p.IPTCDateCreated() == "" {
				c.Fail()
			}

		})

		it(f, "is to EXIF:DateTimeOriginal", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos don't support IPTC
				c.Skip()
				return
			}

			if p.IPTCDateCreated() == "" {
				c.Skip()
				return
			}

			if p.IPTCDateCreated() != strings.Split(p.EXIFDateTimeOriginal(), " ")[0] {
				c.Fail()
			}

		})

	})

	field("IPTC:TimeCreated", func(f string) {

		it(f, "is set", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos don't support IPTC
				c.Skip()
				return
			}

			if p.IPTCTimeCreated() == "" {
				c.Fail()
			}

		})

		it(f, "is set to EXIF:DateTimeOriginal", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideo() || p.IsVideoPreview() {
				// videos don't support IPTC
				c.Skip()
				return
			}

			if p.IPTCTimeCreated() == "" {
				c.Skip()
				return
			}

			if p.IPTCTimeCreated() != strings.Split(p.EXIFDateTimeOriginal(), " ")[1] {
				c.Fail()
			}

		})

	})

	field("QuickTime:CreationDate", func(f string) {

		it(f, "is set to XMP:DateTimeOriginal", func(c *validators.Case, p *picture.Picture) {

			if !p.IsVideo() {
				// photos don't use QuickTime tags
				c.Skip()
				return
			}

			if p.QuickTimeCreationDate() == "" || p.XMPDateTimeOriginal() == "" {
				// This tag is optional (can't be written by exiftool)
				c.Skip()
				return
			}

			if p.QuickTimeCreationDate() != p.XMPDateTimeOriginal() {
				c.Fail()
			}

		})

		it(f, "is set for live videos", func(c *validators.Case, p *picture.Picture) {

			if !p.IsLiveVideo() {
				// only live videos have this tag
				c.Skip()
				return
			}

			if p.QuickTimeCreationDate() == "" {
				c.Fail()
			}

		})

		it(f, "is set to the original picture's EXIF:DateTimeOriginal (with TZ stripped) for live videos", func(c *validators.Case, p *picture.Picture) {

			if !p.IsLiveVideo() {
				// only live videos have this tag
				c.Skip()
				return
			}

			if p.QuickTimeCreationDate()[:19] != p.MasterEXIFDateTimeOriginal() {
				c.Fail()
				return
			}

		})

	})

	for _, f := range []string{"QuickTime:CreateDate", "QuickTime:ModifyDate", "QuickTime:TrackCreateDate", "QuickTime:TrackModifyDate", "QuickTime:MediaCreateDate", "QuickTime:MediaModifyDate"} {

		it(f, "is set", func(c *validators.Case, p *picture.Picture) {

			if !p.IsVideo() {
				// images don't use quicktime tags
				c.Skip()
				return
			}

			if p.ReadStringField(f) == "" {
				c.Fail()
			}

		})

		it(f, "is set to XMP:DateTimeOriginal, in UTC without TZ info", func(c *validators.Case, p *picture.Picture) {

			if !p.IsVideo() {
				// images don't use quicktime tags
				c.Skip()
				return
			}

			if p.ReadStringField(f) == "" || p.XMPDateTimeOriginal() == "" {
				c.Skip()
				return
			}

			t, err := p.ParsedXMPDateTimeOriginal()
			if err != nil {
				c.Fail()
				return
			}

			if p.ReadStringField(f) != t.UTC().Format("2006:01:02 15:04:05") {
				c.Fail()
			}

		})

	}

	field("Filename", func(f string) {

		it(f, "is prefixed by the correct time", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideoPreview() {
				// we don't manage video previews
				c.Skip()
				return
			}

			if p.XMPDateTimeOriginal() == "" {
				c.Skip()
				return
			}

			t, err := p.ParsedXMPDateTimeOriginal()
			if err != nil {
				c.Fail()
				return
			}

			if !strings.HasPrefix(p.Filename(), t.Format("20060102-150405-")) {
				c.Fail()
			}

		})

		it(f, "does not have two time prefixes", func(c *validators.Case, p *picture.Picture) {

			matched, err := regexp.MatchString(`^\d{8}-\d{6}-\d{8}-\d{6}-`, p.XMPDateTimeOriginal())
			if err != nil || matched {
				c.Fail()
			}

		})

		it(f, "is within range of the folder name", func(c *validators.Case, p *picture.Picture) {

			album := p.Album()

			if album == "#misc" {
				c.Skip()
				return
			}

			if len(p.Filename()) < 16 {
				c.Fail()
				return
			}

			parsed, err := time.Parse("20060102-150405-", p.Filename()[0:16])
			if err != nil {
				c.Fail()
				return
			}

			startDate, err := time.Parse("2006.01.02", album[0:10])
			if err != nil {
				c.Fail()
				return
			}
			endDate := startDate

			// day range within month
			matched, err := regexp.MatchString(`^\d{4}.\d{2}.\d{2}-\d{2}-`, album)
			if err != nil {
				c.Fail()
				return
			}
			if matched {

				endDay, err := strconv.Atoi(album[11:13])
				if err != nil {
					c.Fail()
					return
				}

				endDate = time.Date(startDate.Year(), startDate.Month(), endDay, 0, 0, 0, 0, startDate.Location())

			}

			// day range accross months
			matched, err = regexp.MatchString(`^\d{4}.\d{2}.\d{2}-\d{2}.\d{2}-`, album)
			if err != nil {
				c.Fail()
				return
			}
			if matched {

				endMonth, err := strconv.Atoi(album[11:13])
				if err != nil {
					c.Fail()
					return
				}

				endDay, err := strconv.Atoi(album[14:16])
				if err != nil {
					c.Fail()
					return
				}

				endDate = time.Date(startDate.Year(), time.Month(endMonth), endDay, 0, 0, 0, 0, startDate.Location())

			}

			// day range accross years
			matched, err = regexp.MatchString(`^\d{4}.\d{2}.\d{2}-\d{4}.\d{2}.\d{2}-`, album)
			if err != nil {
				c.Fail()
				return
			}
			if matched {

				endYear, err := strconv.Atoi(album[11:15])
				if err != nil {
					c.Fail()
					return
				}

				endMonth, err := strconv.Atoi(album[16:18])
				if err != nil {
					c.Fail()
					return
				}

				endDay, err := strconv.Atoi(album[19:21])
				if err != nil {
					c.Fail()
					return
				}

				endDate = time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, startDate.Location())

			}

			// albums for one day can end at 2PM the next day to account for nights out
			endDate = endDate.Add(24 * time.Hour).Add(14 * time.Hour)

			if parsed.Before(startDate) || parsed.After(endDate) {
				c.Fail()
				return
			}

		})

		it(f, "belongs to year folder", func(c *validators.Case, p *picture.Picture) {

			album := p.YearAlbum()

			if album == "#rush" {
				c.Skip()
				return
			}

			matched, err := regexp.MatchString(`^\d{4}$`, album)
			if err != nil || !matched {
				c.Fail()
				return
			}

			year, err := strconv.Atoi(album)
			if err != nil {
				c.Fail()
				return
			}

			// New years pictures are stored in the ending year
			if !strings.HasPrefix(p.Filename(), album) && !strings.HasPrefix(p.Filename(), strconv.Itoa(year+1)+"0101-") {
				c.Fail()
			}

		})

		it(f, "matches known patterns", func(c *validators.Case, p *picture.Picture) {

			if p.XMPDateTimeOriginal() == "" {
				c.Skip()
				return
			}

			tMetadata, err := p.ParsedXMPDateTimeOriginal()
			if err != nil {
				c.Fail()
				return
			}

			original, err := p.OriginalFilename()
			if err != nil {
				c.Fail()
				return
			}

			original = strings.TrimSuffix(original, filepath.Ext(original))

			formats := []string{
				"20060102 150405",
				"20060102_150405",
				"20060102_150405_HDR",
				"20060102_150405(1)",
				"20060102-150405",
				"20060102150405",
				"Bebop2_20060102150405-0700",
				"ProShot_20060102_150405",
				"VID_20060102_150405",
				"IMG_20060102_150405",
				"PANO_20060102_150405",
				"Photo 02-01-2006 15 04 05",
				"Photo on 02-01-2006 at 15.04",
				"Screen Shot 2006-01-02 at 15.04.05",
				"Screenshot_2006-01-02-15-04-05",
				"signal-2006-01-02-150405",
			}

			for _, f := range formats {

				tOriginal, err := time.Parse(f, original)
				if err != nil {
					// Ignore if we can't match it
					continue
				}

				// formatted filename has no timezone, original can have one, so we compare time values directly
				if tMetadata.Year() != tOriginal.Year() || tMetadata.Month() != tOriginal.Month() || tMetadata.Day() != tOriginal.Day() || tMetadata.Hour() != tOriginal.Hour() || tMetadata.Minute() != tOriginal.Minute() || tMetadata.Second() != tOriginal.Second() {
					c.Fail()
					return
				}

			}

			// This format is UTC time
			formats = []string{
				"2006-01-02 15.04.05",
			}

			for _, f := range formats {

				tOriginal, err := time.Parse(f, original)
				if err != nil {
					// Ignore if we can't match it
					continue
				}

				if tMetadata.Year() != tOriginal.Year() || tMetadata.Month() != tOriginal.Month() || tMetadata.Day() != tOriginal.Day() || tMetadata.Hour() != tOriginal.Hour() || tMetadata.Minute() != tOriginal.Minute() || tMetadata.Second() != tOriginal.Second() {
					// this format can also be UTC time
					tMetadata = tMetadata.UTC()
					if tMetadata.Year() != tOriginal.Year() || tMetadata.Month() != tOriginal.Month() || tMetadata.Day() != tOriginal.Day() || tMetadata.Hour() != tOriginal.Hour() || tMetadata.Minute() != tOriginal.Minute() || tMetadata.Second() != tOriginal.Second() {
						c.Fail()
						return
					}
				}

			}

		})

	})

}

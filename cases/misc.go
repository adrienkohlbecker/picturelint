package cases

import (
	"github.com/adrienkohlbecker/picturelint/picture"
	"github.com/adrienkohlbecker/picturelint/validators"
)

func init() {

	field("File", func(f string) {

		it(f, "should be deleted", func(c *validators.Case, p *picture.Picture) {

			if p.IsVideoPreview() {
				// video previews should be deleted
				c.Fail()
				return
			}

		})

	})

}

package cases

import (
	"fmt"

	"github.com/adrienkohlbecker/picturelint/picture"
	"github.com/adrienkohlbecker/picturelint/validators"
)

type kase struct {
	test   func(v *validators.Case, p *picture.Picture)
	field  string
	legend string
}

var kases = make([]kase, 0)

var currentField string

func field(f string, testSuite func(f string)) {
	testSuite(f)
}

func it(field string, legend string, test func(v *validators.Case, p *picture.Picture)) {

	kases = append(kases, kase{test, field, legend})

}

func Run(p *picture.Picture) []*validators.Case {

	results := make([]*validators.Case, 0)

	for _, k := range kases {

		v := &validators.Case{Legend: fmt.Sprintf("%s %s", k.field, k.legend)}
		k.test(v, p)
		if v.Undefined() {
			v.Success()
		}
		results = append(results, v)

	}

	return results

}

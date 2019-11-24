package cases

import (
	"sync"
	"testing"

	"github.com/adrienkohlbecker/picturelint/picture"
	"github.com/adrienkohlbecker/picturelint/validators"
	"gotest.tools/assert"
)

var _results = make(map[string][]*validators.Case)
var _m = sync.Mutex{}

func CaseResult(t *testing.T, path string, legend string) *validators.Status {

	_m.Lock()
	results, ok := _results[path]

	if !ok {

		p, err := picture.Load(path)
		assert.NilError(t, err)

		results = Run(p)
		_results[path] = results

	}
	_m.Unlock()

	for _, r := range results {
		if r.Legend == legend {
			return &r.Status
		}
	}

	return nil

}

package filter_test

import (
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
	"github.com/suyashkumar/filter"
)

type Thing struct {
	A string
	B int64
}

func TestNewConstraints(t *testing.T) {
	assert := assert.New(t)

	cons, err := filter.NewConstraints(Thing{})

	assert.NoError(err, "There should be no error in a expected call to NewConstraints")
	assert.NotNil(cons, "Constraints should not be nil in happy path")
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	filterStr := "HI"
	tArr := []Thing{Thing{A: filterStr, B: 12}, Thing{A: filterStr, B: 11}, Thing{A: "Not HI", B: 10}}

	cons, _ := filter.NewConstraints(Thing{})

	cons.Add("A", filterStr)

	out, err := filter.Filter(tArr, cons)
	assert.NoError(err, "No error in happy path")

	//TODO(suyashkumar): Stronger, safer test
	for _, e := range out {
		if reflect.DeepEqual(reflect.ValueOf(e), reflect.ValueOf(nil)) {
			continue
		}
		assert.True(
			reflect.DeepEqual(
				reflect.ValueOf(e).FieldByName("A").Interface(),
				reflect.ValueOf(filterStr).Interface(),
			),
			"The filtered field should match the filtered value",
		)
	}
}

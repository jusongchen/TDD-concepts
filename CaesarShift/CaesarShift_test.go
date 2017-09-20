package CaesarShift

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tcase struct {
	source string
	offset int
	code   string
}

func TestBasic(t *testing.T) {
	tc := []tcase{
		tcase{source: "", offset: 0, code: ""},
		tcase{source: "a", offset: 0, code: "a"},
		// testCases{input: nil, offset: 0, expected: "a"},
		tcase{source: "a", offset: 1, code: "b"},
		tcase{source: "ß", offset: 1, code: "ß"},
		tcase{source: "a", offset: 27, code: "b"},
		tcase{source: "a", offset: -1, code: "z"},
		tcase{source: "a", offset: -26 - 1, code: "z"},
		tcase{source: "a", offset: -26*2 - 1, code: "z"},
		tcase{source: "aß", offset: 1, code: "bß"},
		tcase{source: "ßa", offset: 1, code: "ßb"},
		tcase{source: "A", offset: 2, code: "C"},
		tcase{source: "A", offset: -2, code: "Y"},
		tcase{source: "a", offset: 2, code: "c"},
		tcase{source: "z", offset: 1, code: "a"},
		tcase{source: "Z", offset: 2, code: "B"},
		tcase{source: "ab", offset: 2, code: "cd"},
		tcase{source: "a009b", offset: 2, code: "c009d"},
		tcase{source: "a居松b", offset: 2, code: "c居松d"},
		tcase{source: "居松", offset: 2, code: "居松"},
		tcase{source: `abzyABYZ"';:[]091{}()居松`, offset: 1, code: `bcazBCZA"';:[]091{}()居松`},
		tcase{source: `abzyABYZ"';:[]091{}()居松`, offset: -1, code: `zayxZAXY"';:[]091{}()居松`},
	}

	for _, c := range tc {
		actual := Encode(c.source, c.offset)
		assert.Equal(t, c.code, actual, "encode %s with shift position %d", c.source, c.offset)
	}

	for _, c := range tc {
		actual := Decode(c.code, c.offset)
		assert.Equal(t, c.source, actual, "Decode %s with shift position %d", c.code, c.offset)
	}
}

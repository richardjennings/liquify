package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse_Inline(t *testing.T) {
	ip := IncludeParser{}
	v, err := ip.Parse(`some/filepath.whatever arg="test" another='test' yetanother=a.b`)
	expected := "{[{some/filepath.whatever} {arg} {test} {another} {test} {yetanother} {{{a} b}}]}"
	assert.Equal(t, expected, fmt.Sprintf("%s", v))
	assert.Nil(t, err)
}

func TestParse_Multiline(t *testing.T) {
	ip := IncludeParser{}
	v, err := ip.Parse(`some/filepath.whatever 
arg="test" 
another='test' 
yetanother=a.b`)
	expected := "{[{some/filepath.whatever} {arg} {test} {another} {test} {yetanother} {{{a} b}}]}"
	assert.Equal(t, expected, fmt.Sprintf("%s", v))
	assert.Nil(t, err)
}

func TestParse_Multiline_alt(t *testing.T) {
	ip := IncludeParser{}
	v, err := ip.Parse(`some/filepath.whatever 
arg="test"
yetanother=a.b
another='test'`)
	expected := "{[{some/filepath.whatever} {arg} {test} {yetanother} {{{a} b}} {another} {test}]}"
	assert.Equal(t, expected, fmt.Sprintf("%s", v))
	assert.Nil(t, err)
}

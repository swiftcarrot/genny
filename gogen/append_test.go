package gogen

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/swiftcarrot/genny"
)

func Test_Append(t *testing.T) {
	r := require.New(t)

	path := "models/xq.go"
	f := genny.NewFile(path, strings.NewReader(beforeAppend))

	expressions := strings.Split(`
func (xq XQ) Something() {
somethingPrivate()
}`, "\n")

	f, err := Append(f, expressions...)
	r.NoError(err)

	b, err := ioutil.ReadAll(f)
	r.NoError(err)

	r.Equal(path, f.Name())
	r.Equal(afterAppend, string(b))
}

const beforeAppend = `
package models

type XQ struct {
	A string
	W int
}`

const afterAppend = `
package models

type XQ struct {
	A string
	W int
}

func (xq XQ) Something() {
somethingPrivate()
}`

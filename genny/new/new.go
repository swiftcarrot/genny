package new

import (
	"github.com/gobuffalo/flect/name"
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()

	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	if err := g.Box(packr.NewBox("../new/templates")); err != nil {
		return g, errors.WithStack(err)
	}
	name := name.New(opts.Name)
	ctx := plush.NewContext()
	ctx.Set("name", name)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Replace("-name-", name.File().String()))
	g.Transformer(genny.Dot())
	return g, nil
}

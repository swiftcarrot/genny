package attrs

import (
	"strings"

	"github.com/gobuffalo/flect/name"
	"github.com/pkg/errors"
)

func Parse(arg string) (Attr, error) {
	arg = strings.TrimSpace(arg)
	attr := Attr{
		Original:   arg,
		commonType: "string",
	}
	if len(arg) == 0 {
		return attr, errors.New("argument can not be blank")
	}

	parts := strings.Split(arg, ":")
	attr.Name = name.New(parts[0])

	var err error
	var ct string
	if len(parts) > 1 {
		ct = parts[1]
	}
	if attr.commonType, err = UnknownToCommon(ct); err != nil {
		return attr, errors.WithStack(err)
	}

	var gt string
	if len(parts) > 2 {
		gt = parts[2]
	}
	if attr.goType, err = CommonToGo(gt); err != nil {
		return attr, errors.WithStack(err)
	}

	return attr, nil
}

func ParseArgs(args ...string) (Attrs, error) {
	var attrs Attrs

	for _, arg := range args {
		a, err := Parse(arg)
		if err != nil {
			return attrs, errors.WithStack(err)
		}
		attrs = append(attrs, a)
	}

	return attrs, nil
}

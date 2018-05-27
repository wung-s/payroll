package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wung-s/payroll/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}

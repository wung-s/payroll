package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/wung-s/payroll/models"
)

// GroupsList default implementation.
func GroupsList(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	groups := &models.Groups{}
	if err := tx.All(groups); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(groups))
}

package actions

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
	"github.com/wung-s/payroll/models"
)

// RecordsList returns all the records
func RecordsList(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	records := &models.Records{}
	if err := tx.Order("work_date").All(records); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(records))
}

// RecordsUpload default implementation.
func RecordsUpload(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	f, err := c.File("myfile")
	if err != nil {
		return errors.WithStack(err)
	}

	reader := csv.NewReader(f)
	record, err := reader.ReadAll()
	if err != nil {
		return errors.WithStack(err)
	}

	arrLen := len(record)
	reportName := record[arrLen-1][1]

	if arrLen < 2 {
		c.Render(422, r.JSON("invalid file format"))
	}

	exist, err := tx.Where("name = ?", reportName).Exists(&models.Report{})
	if err != nil {
		return errors.WithStack(err)
	}

	if exist {
		c.Render(422, r.JSON("record already exist"))
	}

	report := &models.Report{
		Name: reportName,
	}

	if err = tx.Create(report); err != nil {
		return errors.WithStack(err)
	}

	groups := &models.Groups{}
	if err := tx.All(groups); err != nil {
		return errors.WithStack(err)
	}

	groupsByName := map[string]uuid.UUID{}
	for _, grp := range *groups {
		groupsByName[grp.Name] = grp.ID
	}

	for i, line := range record {
		if i == 0 || i == (arrLen-1) {
			continue
		}

		duration, _ := strconv.ParseFloat(line[1], 64)

		date, err := time.Parse("_2/01/2006 MST", line[0]+ " " + os.Getenv("TIMEZONE"))
		if err != nil {
			return errors.WithStack(err)
		}

		rcrd := &models.Record{
			WorkDate:    date.UTC(),
			DurationHrs: duration,
			Employee:    line[2],
			GroupID:     groupsByName[line[3]],
			ReportID:    report.ID,
		}

		if err := tx.Create(rcrd); err != nil {
			return errors.WithStack(err)
		}
	}

	return c.Render(200, r.JSON("Processed"))
}

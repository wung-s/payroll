package grifts

import (
	"fmt"
	"log"

	"github.com/markbates/grift/grift"
	"github.com/wung-s/payroll/models"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here

		addGroup("A", 20)
		addGroup("B", 30)
		return nil
	})

	grift.Desc("reset", "Truncates all the tables except")
	grift.Add("reset", func(c *grift.Context) error {

		// Note: Any new table created should also be listed here
		sql := `TRUNCATE schema_migration, records, reports, groups CASCADE;`
		if err := models.DB.RawQuery(sql).Exec(); err != nil {
			log.Fatalf("error truncating tables: %v", err)
		} else {
			log.Println("tables truncated successfully")
		}
		return nil
	})

})

func addGroup(name string, rate int) {
	grp := &models.Group{
		Name:       name,
		HourlyRate: rate,
	}

	if err := models.DB.Create(grp); err != nil {
		fmt.Println("could not add group:", err)
	}
}

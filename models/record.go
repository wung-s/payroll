package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type Record struct {
	ID          uuid.UUID `json:"id" db:"id"`
	WorkDate    time.Time `json:"workDate" db:"work_date"`
	DurationHrs float64   `json:"durationHrs" db:"duration_hrs"`
	Employee    string    `json:"employee" db:"employee"`
	GroupID     uuid.UUID `json:"groupID" db:"group_id"`
	ReportID    uuid.UUID `json:"reportID" db:"report_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (r Record) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Records is not required by pop and may be deleted
type Records []Record

// String is not required by pop and may be deleted
func (r Records) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Record) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Record) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Record) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

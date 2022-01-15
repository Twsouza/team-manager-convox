package models

import (
	"strings"
	"time"

	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

const (
	memberTypeInvalid = "The member type must be 'employee' or 'contractor'"
)

var (
	memberTypes = []string{"employee", "contractor"}
)

// Member has a name, type and a tags
type Member struct {
	ID               uuid.UUID     `json:"id" db:"id"`
	CreatedAt        time.Time     `json:"-" db:"created_at"`
	UpdatedAt        time.Time     `json:"-" db:"updated_at"`
	Name             string        `json:"name" db:"name"`
	Type             string        `json:"type" db:"type"`
	ContractDuration int           `json:"contract_duration,omitempty" db:"contract_duration"`
	Tags             slices.String `json:"tags" db:"tags"`
}

// Members is a list of members
type Members []Member

// Validate the member
func (m *Member) Validate(tx *pop.Connection) (*validate.Errors, error) {
	verrs := validate.Validate(
		&validators.StringIsPresent{Name: "Name", Field: m.Name},
		&validators.StringInclusion{Name: "Type", Field: m.Type, List: memberTypes, Message: memberTypeInvalid},
	)

	if m.Type == "employee" && m.ContractDuration != 0 {
		verrs.Add("type", "Employee must not have contract duration")
	}

	if m.Type == "contractor" && m.ContractDuration == 0 {
		verrs.Add("type", "Contractor must have a contract contract")
	}

	return verrs, nil
}

// BeforeSave (create or update), change the tag to lower case
func (m *Member) BeforeSave(tx *pop.Connection) error {
	for i, t := range m.Tags {
		m.Tags[i] = strings.ToLower(t)
	}

	return nil
}

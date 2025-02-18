package user

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/ardanlabs/service/business/sys/validate"
	"github.com/google/uuid"
)

// User represents an individual user.
type User struct {
	ID           uuid.UUID    `json:"id"`
	Name         string       `json:"name"`
	Email        mail.Address `json:"email"`
	Roles        []Role       `json:"roles"`
	PasswordHash []byte       `json:"-"`
	Department   string       `json:"department"`
	Enabled      bool         `json:"enabled"`
	DateCreated  time.Time    `json:"dateCreated"`
	DateUpdated  time.Time    `json:"dateUpdated"`
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name            string       `json:"name" validate:"required"`
	Email           mail.Address `json:"email" validate:"required,email"`
	Roles           []Role       `json:"roles" validate:"required"`
	Department      string       `json:"department"`
	Password        string       `json:"password" validate:"required"`
	PasswordConfirm string       `json:"passwordConfirm" validate:"eqfield=Password"`
}

// Validate checks the data in the model is considered clean.
func (nu NewUser) Validate() error {
	if err := validate.Check(nu); err != nil {
		return err
	}
	return nil
}

// UpdateUser defines what information may be provided to modify an existing
// User. All fields are optional so clients can send just the fields they want
// changed. It uses pointer fields so we can differentiate between a field that
// was not provided and a field that was provided as explicitly blank. Normally
// we do not want to use pointers to basic types but we make exceptions around
// marshalling/unmarshalling.
type UpdateUser struct {
	Name            *string       `json:"name"`
	Email           *mail.Address `json:"email" validate:"omitempty,email"`
	Roles           []Role        `json:"roles"`
	Department      *string       `json:"department"`
	Password        *string       `json:"password"`
	PasswordConfirm *string       `json:"passwordConfirm" validate:"omitempty,eqfield=Password"`
	Enabled         *bool         `json:"enabled"`
}

// Validate checks the data in the model is considered clean.
func (uu UpdateUser) Validate() error {
	if err := validate.Check(uu); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

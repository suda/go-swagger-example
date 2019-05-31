// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Tweet tweet
// swagger:model Tweet
type Tweet struct {

	// contents
	// Required: true
	Contents *string `json:"contents"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty" gorm:"primary_key"`

	// number of characters
	// Minimum: 1
	NumberOfCharacters int64 `json:"numberOfCharacters,omitempty"`

	// user Id
	// Read Only: true
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this tweet
func (m *Tweet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfCharacters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tweet) validateContents(formats strfmt.Registry) error {

	if err := validate.Required("contents", "body", m.Contents); err != nil {
		return err
	}

	return nil
}

func (m *Tweet) validateNumberOfCharacters(formats strfmt.Registry) error {

	if swag.IsZero(m.NumberOfCharacters) { // not required
		return nil
	}

	if err := validate.MinimumInt("numberOfCharacters", "body", int64(m.NumberOfCharacters), 1, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tweet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tweet) UnmarshalBinary(b []byte) error {
	var res Tweet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
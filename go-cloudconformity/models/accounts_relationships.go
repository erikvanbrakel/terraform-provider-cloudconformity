// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountsRelationships accounts relationships
// swagger:model accountsRelationships
type AccountsRelationships struct {

	// organisation
	Organisation *Organisation `json:"organisation,omitempty"`
}

// Validate validates this accounts relationships
func (m *AccountsRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganisation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsRelationships) validateOrganisation(formats strfmt.Registry) error {

	if swag.IsZero(m.Organisation) { // not required
		return nil
	}

	if m.Organisation != nil {
		if err := m.Organisation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organisation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountsRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountsRelationships) UnmarshalBinary(b []byte) error {
	var res AccountsRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

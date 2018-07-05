// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DeleteAccountOKBodyMeta delete account o k body meta
// swagger:model deleteAccountOKBodyMeta
type DeleteAccountOKBodyMeta struct {

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this delete account o k body meta
func (m *DeleteAccountOKBodyMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteAccountOKBodyMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteAccountOKBodyMeta) UnmarshalBinary(b []byte) error {
	var res DeleteAccountOKBodyMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

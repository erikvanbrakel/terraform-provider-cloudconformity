// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AccountAttributesAccessKeys account attributes access keys
// swagger:model accountAttributesAccessKeys
type AccountAttributesAccessKeys struct {

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// role arn
	RoleArn string `json:"roleArn,omitempty"`
}

// Validate validates this account attributes access keys
func (m *AccountAttributesAccessKeys) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountAttributesAccessKeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAttributesAccessKeys) UnmarshalBinary(b []byte) error {
	var res AccountAttributesAccessKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

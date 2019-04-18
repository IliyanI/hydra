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

// Swaggeroauth2TokenParameters Swaggeroauth2TokenParameters swaggeroauth2 token parameters
// swagger:model Swaggeroauth2TokenParameters
type Swaggeroauth2TokenParameters struct {

	// in: formData
	ClientID string `json:"client_id,omitempty"`

	// in: formData
	Code string `json:"code,omitempty"`

	// in: formData
	// Required: true
	GrantType *string `json:"grant_type"`

	// in: formData
	RedirectURI string `json:"redirect_uri,omitempty"`
}

// Validate validates this swaggeroauth2 token parameters
func (m *Swaggeroauth2TokenParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrantType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Swaggeroauth2TokenParameters) validateGrantType(formats strfmt.Registry) error {

	if err := validate.Required("grant_type", "body", m.GrantType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Swaggeroauth2TokenParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Swaggeroauth2TokenParameters) UnmarshalBinary(b []byte) error {
	var res Swaggeroauth2TokenParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
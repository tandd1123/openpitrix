// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGroupDetail openpitrix group detail
// swagger:model openpitrixGroupDetail
type OpenpitrixGroupDetail struct {

	// group base info
	Group *OpenpitrixGroup `json:"group,omitempty"`

	// user set
	UserSet OpenpitrixGroupDetailUserSet `json:"user_set"`
}

// Validate validates this openpitrix group detail
func (m *OpenpitrixGroupDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixGroupDetail) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {

		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGroupDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGroupDetail) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGroupDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
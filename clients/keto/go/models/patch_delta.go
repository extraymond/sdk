// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchDelta patch delta
//
// swagger:model PatchDelta
type PatchDelta struct {

	// action
	Action PatchAction `json:"action,omitempty"`

	// relation tuple
	RelationTuple *InternalRelationTuple `json:"relation_tuple,omitempty"`
}

// Validate validates this patch delta
func (m *PatchDelta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationTuple(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchDelta) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action")
		}
		return err
	}

	return nil
}

func (m *PatchDelta) validateRelationTuple(formats strfmt.Registry) error {

	if swag.IsZero(m.RelationTuple) { // not required
		return nil
	}

	if m.RelationTuple != nil {
		if err := m.RelationTuple.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relation_tuple")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchDelta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchDelta) UnmarshalBinary(b []byte) error {
	var res PatchDelta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

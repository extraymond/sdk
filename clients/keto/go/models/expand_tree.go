// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExpandTree expand tree
//
// swagger:model expandTree
type ExpandTree struct {

	// children
	Children []*ExpandTree `json:"children"`

	// subject
	// Required: true
	Subject Subject `json:"subject"`

	// type
	// Required: true
	// Enum: [union exclusion intersection leaf]
	Type *string `json:"type"`
}

// Validate validates this expand tree
func (m *ExpandTree) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpandTree) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	for i := 0; i < len(m.Children); i++ {
		if swag.IsZero(m.Children[i]) { // not required
			continue
		}

		if m.Children[i] != nil {
			if err := m.Children[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ExpandTree) validateSubject(formats strfmt.Registry) error {

	if err := m.Subject.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subject")
		}
		return err
	}

	return nil
}

var expandTreeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["union","exclusion","intersection","leaf"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		expandTreeTypeTypePropEnum = append(expandTreeTypeTypePropEnum, v)
	}
}

const (

	// ExpandTreeTypeUnion captures enum value "union"
	ExpandTreeTypeUnion string = "union"

	// ExpandTreeTypeExclusion captures enum value "exclusion"
	ExpandTreeTypeExclusion string = "exclusion"

	// ExpandTreeTypeIntersection captures enum value "intersection"
	ExpandTreeTypeIntersection string = "intersection"

	// ExpandTreeTypeLeaf captures enum value "leaf"
	ExpandTreeTypeLeaf string = "leaf"
)

// prop value enum
func (m *ExpandTree) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, expandTreeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExpandTree) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpandTree) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpandTree) UnmarshalBinary(b []byte) error {
	var res ExpandTree
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

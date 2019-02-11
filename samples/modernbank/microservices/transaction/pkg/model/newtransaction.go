// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Newtransaction newtransaction
// swagger:model newtransaction
type Newtransaction struct {

	// amount
	// Required: true
	Amount *float64 `json:"amount"`

	// receiver
	// Required: true
	Receiver *int64 `json:"receiver"`

	// sender
	// Required: true
	Sender *int64 `json:"sender"`
}

// Validate validates this newtransaction
func (m *Newtransaction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceiver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Newtransaction) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *Newtransaction) validateReceiver(formats strfmt.Registry) error {

	if err := validate.Required("receiver", "body", m.Receiver); err != nil {
		return err
	}

	return nil
}

func (m *Newtransaction) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Newtransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Newtransaction) UnmarshalBinary(b []byte) error {
	var res Newtransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
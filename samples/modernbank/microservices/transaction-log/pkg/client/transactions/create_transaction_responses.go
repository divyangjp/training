// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// CreateTransactionReader is a Reader for the CreateTransaction structure.
type CreateTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateTransactionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateTransactionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTransactionCreated creates a CreateTransactionCreated with default headers values
func NewCreateTransactionCreated() *CreateTransactionCreated {
	return &CreateTransactionCreated{}
}

/*CreateTransactionCreated handles this case with default header values.

Created
*/
type CreateTransactionCreated struct {
	Payload *CreateTransactionCreatedBody
}

func (o *CreateTransactionCreated) Error() string {
	return fmt.Sprintf("[POST /transactions][%d] createTransactionCreated  %+v", 201, o.Payload)
}

func (o *CreateTransactionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateTransactionCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionBadRequest creates a CreateTransactionBadRequest with default headers values
func NewCreateTransactionBadRequest() *CreateTransactionBadRequest {
	return &CreateTransactionBadRequest{}
}

/*CreateTransactionBadRequest handles this case with default header values.

Nice try! You can't send negative amounts...
*/
type CreateTransactionBadRequest struct {
}

func (o *CreateTransactionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transactions][%d] createTransactionBadRequest ", 400)
}

func (o *CreateTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTransactionInternalServerError creates a CreateTransactionInternalServerError with default headers values
func NewCreateTransactionInternalServerError() *CreateTransactionInternalServerError {
	return &CreateTransactionInternalServerError{}
}

/*CreateTransactionInternalServerError handles this case with default header values.

Internal server error
*/
type CreateTransactionInternalServerError struct {
}

func (o *CreateTransactionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /transactions][%d] createTransactionInternalServerError ", 500)
}

func (o *CreateTransactionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateTransactionCreatedBody create transaction created body
swagger:model CreateTransactionCreatedBody
*/
type CreateTransactionCreatedBody struct {
	model.Transaction

	model.Version
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *CreateTransactionCreatedBody) UnmarshalJSON(raw []byte) error {
	// CreateTransactionCreatedBodyAO0
	var createTransactionCreatedBodyAO0 model.Transaction
	if err := swag.ReadJSON(raw, &createTransactionCreatedBodyAO0); err != nil {
		return err
	}
	o.Transaction = createTransactionCreatedBodyAO0

	// CreateTransactionCreatedBodyAO1
	var createTransactionCreatedBodyAO1 model.Version
	if err := swag.ReadJSON(raw, &createTransactionCreatedBodyAO1); err != nil {
		return err
	}
	o.Version = createTransactionCreatedBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o CreateTransactionCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	createTransactionCreatedBodyAO0, err := swag.WriteJSON(o.Transaction)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createTransactionCreatedBodyAO0)

	createTransactionCreatedBodyAO1, err := swag.WriteJSON(o.Version)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createTransactionCreatedBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create transaction created body
func (o *CreateTransactionCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with model.Transaction
	if err := o.Transaction.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with model.Version
	if err := o.Version.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *CreateTransactionCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateTransactionCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateTransactionCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

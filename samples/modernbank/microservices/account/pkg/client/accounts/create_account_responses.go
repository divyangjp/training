// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/tetrateio/training/samples/modernbank/microservices/account/pkg/model"
)

// CreateAccountReader is a Reader for the CreateAccount structure.
type CreateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewCreateAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccountCreated creates a CreateAccountCreated with default headers values
func NewCreateAccountCreated() *CreateAccountCreated {
	return &CreateAccountCreated{}
}

/*CreateAccountCreated handles this case with default header values.

Created
*/
type CreateAccountCreated struct {
	/*Version of the microservice that responded
	 */
	Version string

	Payload *model.Account
}

func (o *CreateAccountCreated) Error() string {
	return fmt.Sprintf("[POST /accounts][%d] createAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	o.Payload = new(model.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountInternalServerError creates a CreateAccountInternalServerError with default headers values
func NewCreateAccountInternalServerError() *CreateAccountInternalServerError {
	return &CreateAccountInternalServerError{}
}

/*CreateAccountInternalServerError handles this case with default header values.

Internal server error
*/
type CreateAccountInternalServerError struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *CreateAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /accounts][%d] createAccountInternalServerError ", 500)
}

func (o *CreateAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

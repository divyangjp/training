// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// ListTransactionsReceivedReader is a Reader for the ListTransactionsReceived structure.
type ListTransactionsReceivedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTransactionsReceivedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTransactionsReceivedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListTransactionsReceivedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListTransactionsReceivedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTransactionsReceivedOK creates a ListTransactionsReceivedOK with default headers values
func NewListTransactionsReceivedOK() *ListTransactionsReceivedOK {
	return &ListTransactionsReceivedOK{}
}

/*ListTransactionsReceivedOK handles this case with default header values.

Success!
*/
type ListTransactionsReceivedOK struct {
	/*Version of the microservice that responded
	 */
	Version string

	Payload []*model.Transaction
}

func (o *ListTransactionsReceivedOK) Error() string {
	return fmt.Sprintf("[GET /account/{receiver}/received][%d] listTransactionsReceivedOK  %+v", 200, o.Payload)
}

func (o *ListTransactionsReceivedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTransactionsReceivedNotFound creates a ListTransactionsReceivedNotFound with default headers values
func NewListTransactionsReceivedNotFound() *ListTransactionsReceivedNotFound {
	return &ListTransactionsReceivedNotFound{}
}

/*ListTransactionsReceivedNotFound handles this case with default header values.

No transactions found
*/
type ListTransactionsReceivedNotFound struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *ListTransactionsReceivedNotFound) Error() string {
	return fmt.Sprintf("[GET /account/{receiver}/received][%d] listTransactionsReceivedNotFound ", 404)
}

func (o *ListTransactionsReceivedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

// NewListTransactionsReceivedInternalServerError creates a ListTransactionsReceivedInternalServerError with default headers values
func NewListTransactionsReceivedInternalServerError() *ListTransactionsReceivedInternalServerError {
	return &ListTransactionsReceivedInternalServerError{}
}

/*ListTransactionsReceivedInternalServerError handles this case with default header values.

Internal server error
*/
type ListTransactionsReceivedInternalServerError struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *ListTransactionsReceivedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /account/{receiver}/received][%d] listTransactionsReceivedInternalServerError ", 500)
}

func (o *ListTransactionsReceivedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}

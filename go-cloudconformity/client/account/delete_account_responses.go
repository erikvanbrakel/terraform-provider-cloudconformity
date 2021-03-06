// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikvanbrakel/terraform-provider-cloudconformity/go-cloudconformity/models"
)

// DeleteAccountReader is a Reader for the DeleteAccount structure.
type DeleteAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAccountOK creates a DeleteAccountOK with default headers values
func NewDeleteAccountOK() *DeleteAccountOK {
	return &DeleteAccountOK{}
}

/*DeleteAccountOK handles this case with default header values.

OK
*/
type DeleteAccountOK struct {
	Payload *models.DeleteAccountOKBody
}

func (o *DeleteAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{accountId}][%d] deleteAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAccountAccessSettingsReader is a Reader for the GetAccountAccessSettings structure.
type GetAccountAccessSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountAccessSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountAccessSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountAccessSettingsOK creates a GetAccountAccessSettingsOK with default headers values
func NewGetAccountAccessSettingsOK() *GetAccountAccessSettingsOK {
	return &GetAccountAccessSettingsOK{}
}

/*GetAccountAccessSettingsOK handles this case with default header values.

OK
*/
type GetAccountAccessSettingsOK struct {
}

func (o *GetAccountAccessSettingsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/access][%d] getAccountAccessSettingsOK ", 200)
}

func (o *GetAccountAccessSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetRuleSettingsReader is a Reader for the GetRuleSettings structure.
type GetRuleSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRuleSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRuleSettingsOK creates a GetRuleSettingsOK with default headers values
func NewGetRuleSettingsOK() *GetRuleSettingsOK {
	return &GetRuleSettingsOK{}
}

/*GetRuleSettingsOK handles this case with default header values.

OK
*/
type GetRuleSettingsOK struct {
}

func (o *GetRuleSettingsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/settings/rules][%d] getRuleSettingsOK ", 200)
}

func (o *GetRuleSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
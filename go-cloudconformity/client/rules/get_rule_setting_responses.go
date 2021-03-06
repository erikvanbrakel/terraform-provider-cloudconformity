// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetRuleSettingReader is a Reader for the GetRuleSetting structure.
type GetRuleSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRuleSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRuleSettingOK creates a GetRuleSettingOK with default headers values
func NewGetRuleSettingOK() *GetRuleSettingOK {
	return &GetRuleSettingOK{}
}

/*GetRuleSettingOK handles this case with default header values.

OK
*/
type GetRuleSettingOK struct {
}

func (o *GetRuleSettingOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountId}/settings/rules/{ruleId}][%d] getRuleSettingOK ", 200)
}

func (o *GetRuleSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

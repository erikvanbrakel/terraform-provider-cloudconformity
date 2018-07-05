// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRuleSettingsParams creates a new GetRuleSettingsParams object
// with the default values initialized.
func NewGetRuleSettingsParams() *GetRuleSettingsParams {
	var ()
	return &GetRuleSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuleSettingsParamsWithTimeout creates a new GetRuleSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRuleSettingsParamsWithTimeout(timeout time.Duration) *GetRuleSettingsParams {
	var ()
	return &GetRuleSettingsParams{

		timeout: timeout,
	}
}

// NewGetRuleSettingsParamsWithContext creates a new GetRuleSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRuleSettingsParamsWithContext(ctx context.Context) *GetRuleSettingsParams {
	var ()
	return &GetRuleSettingsParams{

		Context: ctx,
	}
}

// NewGetRuleSettingsParamsWithHTTPClient creates a new GetRuleSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRuleSettingsParamsWithHTTPClient(client *http.Client) *GetRuleSettingsParams {
	var ()
	return &GetRuleSettingsParams{
		HTTPClient: client,
	}
}

/*GetRuleSettingsParams contains all the parameters to send to the API endpoint
for the get rule settings operation typically these are written to a http.Request
*/
type GetRuleSettingsParams struct {

	/*AccountID
	  Account ID

	*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rule settings params
func (o *GetRuleSettingsParams) WithTimeout(timeout time.Duration) *GetRuleSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rule settings params
func (o *GetRuleSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rule settings params
func (o *GetRuleSettingsParams) WithContext(ctx context.Context) *GetRuleSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rule settings params
func (o *GetRuleSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rule settings params
func (o *GetRuleSettingsParams) WithHTTPClient(client *http.Client) *GetRuleSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rule settings params
func (o *GetRuleSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get rule settings params
func (o *GetRuleSettingsParams) WithAccountID(accountID string) *GetRuleSettingsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get rule settings params
func (o *GetRuleSettingsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuleSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
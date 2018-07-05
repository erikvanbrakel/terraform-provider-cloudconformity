// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewGetAccountAccessSettingsParams creates a new GetAccountAccessSettingsParams object
// with the default values initialized.
func NewGetAccountAccessSettingsParams() *GetAccountAccessSettingsParams {
	var ()
	return &GetAccountAccessSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountAccessSettingsParamsWithTimeout creates a new GetAccountAccessSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountAccessSettingsParamsWithTimeout(timeout time.Duration) *GetAccountAccessSettingsParams {
	var ()
	return &GetAccountAccessSettingsParams{

		timeout: timeout,
	}
}

// NewGetAccountAccessSettingsParamsWithContext creates a new GetAccountAccessSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountAccessSettingsParamsWithContext(ctx context.Context) *GetAccountAccessSettingsParams {
	var ()
	return &GetAccountAccessSettingsParams{

		Context: ctx,
	}
}

// NewGetAccountAccessSettingsParamsWithHTTPClient creates a new GetAccountAccessSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountAccessSettingsParamsWithHTTPClient(client *http.Client) *GetAccountAccessSettingsParams {
	var ()
	return &GetAccountAccessSettingsParams{
		HTTPClient: client,
	}
}

/*GetAccountAccessSettingsParams contains all the parameters to send to the API endpoint
for the get account access settings operation typically these are written to a http.Request
*/
type GetAccountAccessSettingsParams struct {

	/*AccountID
	  Account ID

	*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account access settings params
func (o *GetAccountAccessSettingsParams) WithTimeout(timeout time.Duration) *GetAccountAccessSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account access settings params
func (o *GetAccountAccessSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account access settings params
func (o *GetAccountAccessSettingsParams) WithContext(ctx context.Context) *GetAccountAccessSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account access settings params
func (o *GetAccountAccessSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account access settings params
func (o *GetAccountAccessSettingsParams) WithHTTPClient(client *http.Client) *GetAccountAccessSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account access settings params
func (o *GetAccountAccessSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get account access settings params
func (o *GetAccountAccessSettingsParams) WithAccountID(accountID string) *GetAccountAccessSettingsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account access settings params
func (o *GetAccountAccessSettingsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountAccessSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
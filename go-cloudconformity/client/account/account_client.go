// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateAccount creates an account

This endpoint is used to register a new AWS account with Cloud Conformity.
*/
func (a *Client) CreateAccount(params *CreateAccountParams) (*CreateAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAccount",
		Method:             "POST",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAccountOK), nil

}

/*
DeleteAccount deletes account

A DELETE request to this endpoint allows an ADMIN to delete the specified account.
*/
func (a *Client) DeleteAccount(params *DeleteAccountParams) (*DeleteAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAccount",
		Method:             "DELETE",
		PathPattern:        "/accounts/{accountId}",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAccountOK), nil

}

/*
GetAccount gets account details

This endpoint allows you to get the details of the specified account.
*/
func (a *Client) GetAccount(params *GetAccountParams) (*GetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccount",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountOK), nil

}

/*
GetAccountAccessSettings gets account access setting

This endpoint allows ADMIN users to get the current setting Cloud Conformity uses to access the specified account.
*/
func (a *Client) GetAccountAccessSettings(params *GetAccountAccessSettingsParams) (*GetAccountAccessSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountAccessSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccountAccessSettings",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}/access",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountAccessSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountAccessSettingsOK), nil

}

/*
ListAccounts lists all accounts

This endpoint allows you to query all accounts that you have access to.
*/
func (a *Client) ListAccounts(params *ListAccountsParams) (*ListAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAccounts",
		Method:             "GET",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAccountsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAccountsOK), nil

}

/*
ScanAccount scans account

This endpoint allows you to run conformity bot for the specified account.
*/
func (a *Client) ScanAccount(params *ScanAccountParams) (*ScanAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScanAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "scanAccount",
		Method:             "POST",
		PathPattern:        "/accounts/{accountId}/scan",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScanAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ScanAccountOK), nil

}

/*
UpdateAccount updates account

A PATCH request to this endpoint allows changes to the account name, enviornment, and code.
*/
func (a *Client) UpdateAccount(params *UpdateAccountParams) (*UpdateAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAccount",
		Method:             "PATCH",
		PathPattern:        "/accounts/{accountId}",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAccountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAccountOK), nil

}

/*
UpdateAccountSubscription updates account subscription

A PATCH request to this endpoint allows you to change the add-on package subscription of the specified account.
*/
func (a *Client) UpdateAccountSubscription(params *UpdateAccountSubscriptionParams) (*UpdateAccountSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAccountSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAccountSubscription",
		Method:             "PATCH",
		PathPattern:        "/accounts/{id}/subscription",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAccountSubscriptionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAccountSubscriptionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
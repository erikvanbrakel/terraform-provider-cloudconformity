// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRuleSetting gets rule setting

 GET request to this endpoint allows you to get configured rule setting for the specified rule Id of the specified account.
*/
func (a *Client) GetRuleSetting(params *GetRuleSettingParams) (*GetRuleSettingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleSettingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuleSetting",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}/settings/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRuleSettingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRuleSettingOK), nil

}

/*
GetRuleSettings gets rule settings

A GET request to this endpoint allows you to get rule settings for all configured rules of the specified account.
*/
func (a *Client) GetRuleSettings(params *GetRuleSettingsParams) (*GetRuleSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuleSettings",
		Method:             "GET",
		PathPattern:        "/accounts/{accountId}/settings/rules",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRuleSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRuleSettingsOK), nil

}

/*
UpdateRuleSetting updates rule setting

A PATCH request to this endpoint allows you to customize rule setting for the specified rule Id of the specified account.
*/
func (a *Client) UpdateRuleSetting(params *UpdateRuleSettingParams) (*UpdateRuleSettingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRuleSettingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRuleSetting",
		Method:             "PATCH",
		PathPattern:        "/accounts/{accountId}/settings/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRuleSettingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRuleSettingOK), nil

}

/*
UpdateRuleSettings updates rule settings

A PATCH request to this endpoint allows you to customize rule settings for the specified account.
*/
func (a *Client) UpdateRuleSettings(params *UpdateRuleSettingsParams) (*UpdateRuleSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRuleSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRuleSettings",
		Method:             "PATCH",
		PathPattern:        "/accounts/{accountId}/settings/rules",
		ProducesMediaTypes: []string{"application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRuleSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRuleSettingsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

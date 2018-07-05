package provider

import (
	"errors"
	"github.com/erikvanbrakel/terraform-provider-cloudconformity/go-cloudconformity/client"
	"github.com/erikvanbrakel/terraform-provider-cloudconformity/go-cloudconformity/client/account"
	"github.com/erikvanbrakel/terraform-provider-cloudconformity/go-cloudconformity/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCloudConformityAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudConformityAccountCreate,
		Update: resourceCloudConformityAccountUpdate,
		Read:   resourceCloudConformityAccountRead,
		Delete: resourceCloudConformityAccountDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_cost_package": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_security_package": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCloudConformityAccountCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*client.CloudConformity)

	cp := d.Get("enable_cost_package").(bool)
	sp := d.Get("enable_security_package").(bool)

	accountBody := &models.CreateAccountParamsBody{
		Data: &models.Account{
			Type: "account",
			Attributes: &models.AccountAttributes{
				Name:            d.Get("name").(string),
				Environment:     d.Get("environment").(string),
				SecurityPackage: &sp,
				CostPackage:     &cp,
				Access: &models.AccountAttributesAccess{
					Keys: &models.AccountAttributesAccessKeys{
						ExternalID: d.Get("external_id").(string),
						RoleArn:    d.Get("role_arn").(string),
					},
				},
			},
		},
	}

	response, err := c.Account.CreateAccount(account.NewCreateAccountParams().WithAccount(accountBody))

	if err != nil {
		switch v := err.(type) {
		case *account.CreateAccountUnprocessableEntity:
			return errors.New(v.Payload.Errors[0].Detail)
		default:
			return err
		}
	}

	d.SetId(response.Payload.Data.ID)

	return nil
}

func resourceCloudConformityAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	return errors.New("not implemented")
}

func resourceCloudConformityAccountRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*client.CloudConformity)

	response, err := c.Account.GetAccount(account.NewGetAccountParams().WithAccountID(d.Id()))

	if err != nil {
		return err
	}

	data := response.Payload.Data
	d.Set("name", data.Attributes.Name)
	d.Set("environment", data.Attributes.Environment)
	d.Set("enable_cost_package", data.Attributes.CostPackage)
	d.Set("enable_security_package", data.Attributes.SecurityPackage)

	return nil
}

func resourceCloudConformityAccountDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*client.CloudConformity)
	params := account.NewDeleteAccountParams()
	params.SetAccountID(d.Id())

	_, err := c.Account.DeleteAccount(params)

	if err != nil {
		switch err.(type) {
		default:
			return err
		}
	}

	return nil
}

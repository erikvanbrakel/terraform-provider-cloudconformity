package provider

import (
	"github.com/erikvanbrakel/terraform-provider-cloudconformity/go-cloudconformity/client"
	"github.com/erikvanbrakel/terraform-provider-cloudconformity/go-cloudconformity/client/external_id"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCloudConformityExternalId() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudConformityExternalIdRead,

		Schema: map[string]*schema.Schema{},
	}
}

func dataSourceCloudConformityExternalIdRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*client.CloudConformity)

	response, err := c.ExternalID.GetExternalID(external_id.NewGetExternalIDParams())

	if err != nil {
		return err
	}

	d.SetId(response.Payload.Data.ID)

	return nil
}

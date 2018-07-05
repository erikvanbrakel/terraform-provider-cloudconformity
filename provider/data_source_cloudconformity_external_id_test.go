package provider

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccCloudConformityExternalId(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudConformityExternalIdConfig,
			},
		}})
}

var testAccCloudConformityExternalIdConfig = `data "cloudconformity_external_id" "exid" {}`

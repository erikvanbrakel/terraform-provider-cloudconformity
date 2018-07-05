package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

func TestAccCloudConformityAccount(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudConformityAccountConfig,
			},
		}})
}

var externalId = os.Getenv("cc_test_external_id")
var roleArn = os.Getenv("cc_test_role_arn")

var testAccCloudConformityAccountConfig = fmt.Sprintf(`
resource "cloudconformity_account" "account" {
  name = "Test-account"
  environment = "Test-env"

  external_id = "%s"
  role_arn    = "%s"
}
`, externalId, roleArn)

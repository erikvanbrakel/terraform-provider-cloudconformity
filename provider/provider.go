package provider

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"fmt"
	"github.com/erikvanbrakel/terraform-provider-cloudconformity/go-cloudconformity/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"os"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  os.Getenv("CLOUD_CONFORMITY_API_KEY"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudconformity_account": resourceCloudConformityAccount(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloudconformity_external_id": dataSourceCloudConformityExternalId(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	rt := httptransport.New(
		client.DefaultHost,
		client.DefaultBasePath,
		client.DefaultSchemes,
	)
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()
	rt.Producers["application/vnd.api+json"] = runtime.JSONProducer()

	c := client.New(rt, strfmt.Default)

	c.SetTransport(&KeyAuthTransport{ApiKey: d.Get("api_key").(string), Transport: c.Transport})
	return c, nil
}

type KeyAuthTransport struct {
	ApiKey    string
	Transport runtime.ClientTransport
}

func (t *KeyAuthTransport) Submit(o *runtime.ClientOperation) (interface{}, error) {
	o.AuthInfo = &KeyAuth{t.ApiKey}
	return t.Transport.Submit(o)
}

type KeyAuth struct {
	key string
}

func (a *KeyAuth) AuthenticateRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetHeaderParam("Authorization", fmt.Sprintf("ApiKey %s", a.key))
	r.SetHeaderParam("Content-Type", "application/vnd.api+json")
	return nil
}

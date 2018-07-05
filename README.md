# terraform-provider-cloudconformity

Terraform provider for https://www.cloudconformity.com

## Building the provider
In this section you will learn how to build and tun `terraform-provider-cloudconformity` locally. Please follow the
steps below:

Requirements
---------------------

 - [Terraform] 0.10.x
 - [Go] 1.9 (to build the provider plugin)
 - [Cloud Conformity] Must have an account to be able to get an API key


Setting up your environment
---------------------

Create the environment variables below:

`$GOROOT=$HOME/go`

`$GOPATH=$GOROOT/bin`

The `GOPATH` can be set wherever you want but please read this [topic][0] to understand how they work.

Clone repository to: `$GOPATH/src/github.com/erikvanbrakel/terraform-provider-cloudconformity`


## Installation
Download the binary for your platform and architecture from the [releases page]. Unpack the zip, and place the
`terraform-provider-cloudconformity` binary in the same directory as the `terraform` binary or add a `.terraformrc` with
the provider stanza:

```hcl
providers {
    cloudconformity = "/PATH/TO/MODULE/ARCH/terraform-provider-cloudconformity"
}
```


## Usage

```hcl
data "cloudconformity_external_id" "exid" {}

resource "aws_cloudformation_stack" "cloud-conformity" {

  name         = "CloudConformity"

  template_url = "https://s3-us-west-2.amazonaws.com/cloudconformity/CloudConformity.template"

  parameters {
    AccountId  = "717210094962"
    ExternalId = "${data.cloudconformity_external_id.exid.id}"
  }

  capabilities = ["CAPABILITY_NAMED_IAM"]

}

resource "cloudconformity_account" "account" {
  name        = "My Account"
  environment = "My Environment"

  external_id = "${data.cloudconformity_external_id.exid.id}"
  role_arn    = "${aws_cloudformation_stack.cloud-conformity.outputs["CloudConformityRoleArn"]}"
}
```

For more information, please see the [documentation].

[Terraform]: https://www.terraform.io/downloads.html
[Go]: https://golang.org/doc/install
[Cloud Conformity]: https://www.cloudconformity.com
[releases page]: https://github.com/erikvanbrakel/terraform-provider-cloudconformity/releases
[documentation]: docs/README.md
[0]: https://stackoverflow.com/questions/7970390/what-should-be-the-values-of-gopath-and-goroot
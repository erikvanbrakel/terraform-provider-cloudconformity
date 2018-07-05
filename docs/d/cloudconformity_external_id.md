# cloudconformity_external_id

Provides an external ID for your Cloud Conformity organisation.

## Example Usage
```hcl
resource "cloudconformity_external_id" "exid" { }

resource "cloudconformity_account" "my_account" {
    name        = "My Account"
    environment = "My Environment"

    external_id = "${data.cloudconformity_external_id.exid.id}"
    role_arn    = "..."
}
```

## Attributes reference
 - `id` - The external ID

[Back to Index](../README.md)
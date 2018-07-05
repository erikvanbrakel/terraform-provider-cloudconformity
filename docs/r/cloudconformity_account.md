# cloudconformity_account
Provides a Cloud Conformity Account.

## Example Usage
```hcl
resource "cloudconformity_account" "my_account" {
    name        = "My Account"
    environment = "My Environment"

    external_id = "..."
    role_arn    = "..."
}
```

## Argument reference
 - `name` - (required) The name of your account.
 - `environment` - (required) The environment for your account.
 - `external_id` - (required) The external ID for your Cloud Conformity organisation.
 - `role_arn` - (required) The ARN of the role your account can assume.
 - `enable_cost_package` - (optional) Specify if the cost optimization package should be enabled.
 - `enable_security_package` - (optional) Specify if the security package should be enabled.


[Back to Index](../README.md)
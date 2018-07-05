# Cloud Conformity Provider

This provider is used to manage your Cloud Conformity setup. The provider needs to be configured with the proper
credentials before it can be used.

## Example Usage
```hcl
# Configure the Cloud Conformity Provider
provider "cloudconformity" {
    api_key   = "${var.cloudconformity_api_key}"
}
```

## Authentication
The provider offers a flexible means of providing credentials for
authentication. The following methods are supported and explained below:

 - Static credentials
 - Environment variables

### Static credentials
Static credentials can be provided by adding an `api_key` in-line in the provider block:

Usage:
```hcl
provider "cloudconformity" {
    api_key  = "your-api-key"
}
```

### Environment variables
You can provide your credentials via the `CLOUDCONFORMITY_API_KEY` environment variable ,
representing your Cloud Conformity API Key.

Usage:
```hcl
provider "cloudconformity" { }
```

```bash
$ export CLOUDCONFORMITY_API_KEY="your-api-key"
$ terraform plan
```

## Argument Reference
 - `api_key` - (Optional) This is the Cloud Conformity API Key. It must be provided, but
   it can also be sourced from the `CLOUDCONFORMITY_API_KEY` environment variable.

[Back to Index](README.md)

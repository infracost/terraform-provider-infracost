# infracost_aws_dx_gateway_association

Provides estimated usage data for an AWS DX Gateway association.

## Example Usage

```hcl
resource "aws_dx_gateway_association" "gateway" {
    dx_gateway_id = "dx-123456"
    associated_gateway_id = "tgw-123456"
}

data "infracost_aws_dx_gateway_association" "gateway" {
  resources = list(aws_dx_gateway_association.gateway.id)

  monthly_data_processed_gb {
    value = 100
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the DX gateway association(s) to apply the estimated usage.
* `monthly_data_processed_gb` - (Optional) The estimated data processed by the DX gateway association per month in GB.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


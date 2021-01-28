# infracost_aws_nat_gateway

Provides estimated usage data for an AWS Nat Gateway.

## Example Usage

```hcl
resource "aws_nat_gateway" "my_nat_gateway" {
  allocation_id = "eip-12345678"
  subnet_id     = "subnet-12345678"
}

data "infracost_aws_nat_gateway" "my_nat_gateway" {
  resources = list(aws_nat_gateway.my_nat_gateway.id)

  monthly_data_processed_gb {
    value = 10
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the Nat Gateways to apply the estimated usage.
* `monthly_data_processed_gb` - (Optional) The estimated data processed by the NAT Gateway per month in GB.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


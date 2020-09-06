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

  monthly_gb_data_processed {
    value = 10
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the Nat Gateways to apply the estimated usage.
* `monthly_gb_data_processed` - (Optional) The estimated GB of data processed by the NAT Gateway per month. See [monthly_gb_data_processed](#monthly_gb_data_processed) below for details on attributes.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


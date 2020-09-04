# infracost_aws_nat_gateway

Provides estimated usage data for an AWS Nat Gateway.

## Example Usage

```hcl
resource "aws_nat_gateway" "my-nat-gateway" {
  allocation_id = "eip-12345678"
  subnet_id     = "subnet-12345678"
}

resource "infracost_aws_nat_gateway" "my-nat-gateway" {
  resources = list(aws_nat_gateway.my-nat-gateway-1.id)

  gb_data_processed_monthly {
    value = 10
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the Nat Gateways to apply the estimated usage.
* `gb_data_processed_monthly` - (Optional) The estimated GB of data processed by the NAT Gateway per month. See [gb_data_processed_monthly](#gb_data_processed_monthly) below for details on attributes.

### gb_data_processed_monthly

The `gb_data_processed_monthly` supports the following:

* `value` - (Optional) The estimated value.

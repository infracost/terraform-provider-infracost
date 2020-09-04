# Infracost Provider

The Infracost provider is used to add estimated usage of Terraform resources to add better cost estimations when used with the [Infracost](https://infracost.io) tool.

## Example Usage

```hcl
# Configure the Infracost provider
provider "infracost" {}

# Create an AWS NAT Gateway
resource "aws_nat_gateway" "my-nat-gateway" {
  allocation_id = "eip-12345678"
  subnet_id     = "subnet-12345678"
}

# Add usage estimates for the AWS NAT Gateway
resource "infracost_aws_nat_gateway" "my-nat-gateway" {
  resources = list(aws_nat_gateway.my-nat-gateway-1.id)

  gb_data_processed_monthly {
    value = 10
  }
}
```

## Argument Reference

The Infracost provider takes no arguments.

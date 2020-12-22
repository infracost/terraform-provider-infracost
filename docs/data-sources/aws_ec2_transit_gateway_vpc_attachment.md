# infracost_aws_ec2_transit_gateway_vpc_attachment

Provides estimated usage data for an AWS EC2 Transit Gateway attachment.

## Example Usage

```hcl
resource "aws_ec2_transit_gateway_vpc_attachment" "vpc_attachment" {
    subnet_ids = ["subnet-123456", "subnet-654321"]
    transit_gateway_id = "tgw-123456"
    vpc_id = "vpc-123456"
}

data "infracost_aws_ec2_transit_gateway_vpc_attachment" "vpc_attachment" {
  resources = list(aws_ec2_transit_gateway_vpc_attachment.vpc_attachment.id)

  monthly_gb_data_processed {
    value = 100
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the EC2 transit gateway attachment(s) to apply the estimated usage.
* `monthly_gb_data_processed` - (Optional) The estimated GB of data processed by the EC2 transit gateway attachment(s) per month.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


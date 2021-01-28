# infracost_aws_vpn_connection

Provides estimated data processing costs for an AWS VPN Transit gateway attachment.

## Example Usage

```hcl
resource "aws_vpn_connection" "transit" {
  customer_gateway_id = "my-cgw-id"
  transit_gateway_id = "my-tgw-id"
  type = "ipsec.1"
}

data "infracost_aws_vpn_connection" "vpn_transit_costs" {
  resources = list(aws_vpn_connection.transit.id)

  monthly_data_processed_gb {
    value = 100
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the VPN connections to apply the estimated usage.
* `monthly_data_processed_gb` - (Optional) The estimated monthly data processed through a transit gateway attached to your VPN Connection in GB.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.


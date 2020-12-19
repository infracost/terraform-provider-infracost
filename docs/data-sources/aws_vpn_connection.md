# infracost_aws_vpn_connection

Provides estimated usage data for an AWS VPN Transit gateway data processing.

## Example Usage

```hcl
resource "aws_vpn_connection" "transit" {
  customer_gateway_id = "my-cgw-id"
  transit_gateway_id = "my-tgw-id"
  type = "ipsec.1"
}

data "infracost_aws_vpn_connection" "vpn_transit_costs" {
  resources = list(aws_vpn_connection.transit.id)

  monthly_gb_data_processed {
    value = 100
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the VPN connections to apply the estimated usage.
* `monthly_gb_data_processed` - (Optional) The estimated monthly data processed through a transit gateway attached to your VPN Connection.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.


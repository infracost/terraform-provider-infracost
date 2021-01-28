# infracost_aws_cloudwatch_log_group

Provides estimated usage data for an AWS CloudWatch log group.

## Example Usage

```hcl
resource "aws_cloudwatch_log_group" "logs" {
  name = "cloudwatch-log-group"
}

data "infracost_aws_cloudwatch_log_group" "logs" {
  resources = list(aws_cloudwatch_log_group.logs.id)

  monthly_data_ingested_gb {
    value = 1000
  }

  storage_gb {
    value = 1000
  }

  monthly_data_scanned_gb {
    value = 200
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the CloudWatch logs log group to apply the estimated usage.
* `monthly_data_ingested_gb` - (Optional) The estimated data ingested by CloudWatch logs per month in GB.
* `storage_gb` - (Optional) The estimated data stored by CloudWatch logs per month in GB.
* `monthly_data_scanned_gb` - (Optional) The estimated data scanned by CloudWatch logs insights per month in GB.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


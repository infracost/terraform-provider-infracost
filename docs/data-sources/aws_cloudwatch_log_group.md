# infracost_aws_cloudwatch_log_group

Provides estimated usage data for an AWS Cloudwatch log group.

## Example Usage

```hcl
resource "aws_cloudwatch_log_group" "logs" {
  name = "cloudwatch-log-group"
}

data "infracost_aws_cloudwatch_log_group" "logs" {
  resources = list(aws_cloudwatch_log_group.logs.id)

  monthly_gb_data_ingestion {
    value = 1000
  }

  monthly_gb_data_storage {
    value = 1000
  }
  
  monthly_gb_data_scanned {
    value = 200
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the Cloudwatch logs log group to apply the estimated usage.
* `monthly_gb_data_ingestion` - (Optional) The estimated GB of data ingested by Cloudwatch logs per month. 
* `monthly_gb_data_storage` - (Optional) The estimated GB of data stored by Cloudwatch logs per month.
* `monthly_gb_data_scanned` - (Optional) The estimated GB of data scanned by Cloudwatch logs insights per month.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


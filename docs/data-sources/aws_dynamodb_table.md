# infracost_aws_dynamodb_table

Provides estimated usage data for an AWS DynamoDB.

## Example Usage

```hcl
resource "aws_dynamodb_table" "my_dynamodb_table" {
  name           = "GameScores"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "UserId"
  range_key      = "GameTitle"

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "GameTitle"
    type = "S"
  }

  replica {
    region_name = "us-east-2"
  }

  replica {
    region_name = "us-west-1"
  }

}

data "infracost_aws_dynamodb_table" "my_dynamodb_table" {
  resources = list(aws_dynamodb_table.my_dynamodb_table.id)

  monthly_write_request_units {
    value = 3000000
  }
  monthly_read_request_units {
    value = 8000000
  }
  storage_gb {
    value = 230
  }
  pitr_backup_storage_gb {
    value = 2300
  }
  on_demand_backup_storage_gb {
    value = 460
  }
  monthly_data_restored_gb {
    value = 230
  }
  monthly_streams_read_request_units {
    value = 2
  }

}
```

## Argument Reference

* `resources` - (Required) The IDs of the DynamoDBs to apply the estimated usage.
* `monthly_write_request_units` - (Optional) The estimated write request units per month in (used for on-demand DynamoDB).
* `monthly_read_request_units` - (Optional) The estimated read request units per month in (used for on-demand DynamoDB).
* `storage_gb` - (Optional) The estimated total storage for tables per month in GBs.
* `pitr_backup_storage_gb` - (Optional) The estimated total storage for continuous backups (PITR) per month in GBs.
* `on_demand_backup_storage_gb` - (Optional) The estimated total storage for on-demand backups per month in GBs.
* `monthly_data_restored_gb` - (Optional) The estimated size of restored data per month in GBs.
* `monthly_streams_read_request_units` - (Optional) The estimated streams read request units per month.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


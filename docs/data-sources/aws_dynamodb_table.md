# infracost_aws_dynamodb_table

Provides estimated usage data for an AWS DynamoDB.

## Example Usage

```hcl
resource "aws_dynamodb_table" "my-dynamodb-table" {
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

resource "infracost_aws_dynamodb_table" "my-dynamodb-table" {
  resources = list(aws_dynamodb_table.my-dynamodb-table.id)

  monthly_million_write_request_units {
    value = 3
  }
  monthly_million_read_request_units {
    value = 8
  }
  monthly_gb_data_storage {
    value = 230
  }

}
```

## Argument Reference

* `resources` - (Required) The IDs of the DynamoDBs to apply the estimated usage.
* `monthly_million_write_request_units` - (Optional) The estimated write request units per month in millions (used for on-demand DynamoDB).
* `monthly_million_write_request_units` - (Optional) The estimated read request units per month in millions (used for on-demand DynamoDB).
* `monthly_gb_data_storage` - (Optional) The estimated storage for tables per month in GBs.
* `monthly_gb_continuos_backup_storage` - (Optional) The estimated storage for continuos backups(PITR) in GBs.
* `monthly_gb_on_demand_backup_storage` - (Optional) The estimated storage for on-demand backups in GBs.
* `gb_restore` - (Optional) The estimated size of restored data in GBs.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


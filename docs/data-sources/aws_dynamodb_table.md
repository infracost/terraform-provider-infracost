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

data "infracost_aws_dynamodb_table" "my_dynamodb_table" {
  resources = list(aws_dynamodb_table.my_dynamodb_table.id)

  monthly_million_write_request_units {
    value = 3
  }
  monthly_million_read_request_units {
    value = 8
  }
  monthly_gb_data_storage {
    value = 230
  }
  monthly_gb_continuos_backup_storage {
    value = 2300
  }
  monthly_gb_on_demand_backup_storage {
    value = 460
  }
  monthly_gb_restore {
    value = 230
  }
  monthly_gb_data_in {
    value = 10
  }
  monthly_gb_data_out {
    value = 30
  }
  monthly_streams_read_request_units {
    value = 2
  }

}
```

## Argument Reference

* `resources` - (Required) The IDs of the DynamoDBs to apply the estimated usage.
* `monthly_million_write_request_units` - (Optional) The estimated write request units per month in millions (used for on-demand DynamoDB).
* `monthly_million_write_request_units` - (Optional) The estimated read request units per month in millions (used for on-demand DynamoDB).
* `monthly_gb_data_storage` - (Optional) The estimated storage for tables per month in GBs.
* `monthly_gb_continuous_backup_storage` - (Optional) The estimated storage for continuous backups (PITR) per month in GBs.
* `monthly_gb_on_demand_backup_storage` - (Optional) The estimated storage for on-demand backups per month in GBs.
* `monthly_gb_restore` - (Optional) The estimated size of restored data per month in GBs.
* `monthly_gb_data_in` - (Optional) The estimated size of transferred data into DynamoDB per month in GBs.
* `monthly_gb_data_out` - (Optional) The estimated size of transferred data out of DynamoDB per month in GBs.
* `monthly_streams_read_request_units` - (Optional) The estimated streams read request units per month.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.


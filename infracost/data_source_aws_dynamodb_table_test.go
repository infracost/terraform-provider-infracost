package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsDynamoDBTable(t *testing.T) {
	name := "data.infracost_aws_dynamodb_table.my_dynamodb_table_usage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsDynamoDBTableConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceAttrValue(name, "monthly_write_request_units", 3000000.0),
					testCheckResourceAttrValue(name, "monthly_read_request_units", 8000000.0),
					testCheckResourceAttrValue(name, "storage_gb", 230.0),
					testCheckResourceAttrValue(name, "pitr_backup_storage_gb", 2300.0),
					testCheckResourceAttrValue(name, "on_demand_backup_storage_gb", 460.0),
					testCheckResourceAttrValue(name, "monthly_data_restored_gb", 230.0),
					testCheckResourceAttrValue(name, "monthly_streams_read_request_units", 2.0),
				),
			},
		},
	})
}

func testAwsDynamoDBTableConfig() string {
	return `
		data "infracost_aws_dynamodb_table" "my_dynamodb_table_usage" {
			resources = list("my_dynamodb_table")
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
	`
}

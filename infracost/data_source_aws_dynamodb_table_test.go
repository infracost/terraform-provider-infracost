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
					testCheckResourceAttrValue(name, "monthly_gb_data_storage", 230.0),
					testCheckResourceAttrValue(name, "monthly_gb_continuous_backup_storage", 2300.0),
					testCheckResourceAttrValue(name, "monthly_gb_on_demand_backup_storage", 460.0),
					testCheckResourceAttrValue(name, "monthly_gb_restore", 230.0),
					testCheckResourceAttrValue(name, "monthly_gb_data_in", 10.0),
					testCheckResourceAttrValue(name, "monthly_gb_data_out", 30.0),
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
			monthly_gb_data_storage {
				value = 230
			}
			monthly_gb_continuous_backup_storage {
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
	`
}

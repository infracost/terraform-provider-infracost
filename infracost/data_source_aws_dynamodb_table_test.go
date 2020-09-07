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
					testCheckResourceAttrValue(name, "monthly_million_write_request_units", 3.0),
					testCheckResourceAttrValue(name, "monthly_million_read_request_units", 8.0),
					testCheckResourceAttrValue(name, "monthly_gb_data_storage", 230.0),
					testCheckResourceAttrValue(name, "monthly_gb_continuos_backup_storage", 2300.0),
					testCheckResourceAttrValue(name, "monthly_gb_continuos_backup_storage", 470.0),
					testCheckResourceAttrValue(name, "gb_restore", 230.0),
				),
			},
		},
	})
}

func testAwsDynamoDBTableConfig() string {
	return `
		data "infracost_aws_dynamodb_table" "my_dynamodb_table_usage" {
			resources = list("my_dynamodb_table")
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
				value: 2300
			}
			monthly_gb_on_demand_backup_storage {
				value: 460
			}
			gb_restore {
				value: 230
			}
		}
	`
}

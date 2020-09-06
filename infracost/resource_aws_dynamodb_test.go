package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestResourceAwsDynamoDB(t *testing.T) {
	name := "infracost_aws_dynamodb_table.my_dynamodb_table_usage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getTestAwsDynamoDBConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceAttrValue(name, "monthly_million_write_request_units", 3.0),
					testCheckResourceAttrValue(name, "monthly_million_read_request_units", 8.0),
				),
			},
		},
	})
}

func getTestAwsDynamoDBConfig() string {
	return `
		resource "infracost_aws_dynamodb_table" "my_dynamodb_table_usage" {
			resources = list("my_dynamodb_table")

			monthly_million_write_request_units {
				value = 3
			}

			monthly_million_read_request_units {
				value = 8
			}
		}
	`
}

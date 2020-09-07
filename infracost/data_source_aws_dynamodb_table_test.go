package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsDynamoDBTable(t *testing.T) {
	name := "data.infracost_aws_nat_gateway.my_nat_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsDynamoDBTableConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckResourceAttrValue(name, "monthly_million_write_request_units", 3.0),
					testCheckResourceAttrValue(name, "monthly_million_read_request_units", 8.0),
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
		}
	`
}

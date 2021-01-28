package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsCloudwatchLogGroup(t *testing.T) {
	name := "data.infracost_aws_cloudwatch_log_group.log_group"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsCloudwatchLogGroupConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_data_ingested_gb", 1000),
					testCheckResourceAttrValue(name, "storage_gb", 500),
					testCheckResourceAttrValue(name, "monthly_data_scanned_gb", 100),
				),
			},
		},
	})
}

func testAwsCloudwatchLogGroupConfig() string {
	return `
		data "infracost_aws_cloudwatch_log_group" "log_group" {
			resources = list("log_group_1", "log_group_2")

			monthly_data_ingested_gb {
				value = 1000
			}

			storage_gb {
				value = 500
			}

			monthly_data_scanned_gb {
				value = 100
			}
		}
	`
}

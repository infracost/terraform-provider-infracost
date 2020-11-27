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
					testCheckResourceAttrValue(name, "monthly_gb_data_ingestion", 1000),
					testCheckResourceAttrValue(name, "monthly_gb_data_storage", 500),
					testCheckResourceAttrValue(name, "monthly_gb_data_scanned", 100),
				),
			},
		},
	})
}

func testAwsCloudwatchLogGroupConfig() string {
	return `
		data "infracost_aws_cloudwatch_log_group" "log_group" {
			resources = list("log_group_1", "log_group_2")

			monthly_gb_data_ingestion {
				value = 1000
			}
			
			monthly_gb_data_storage {
				value = 500
			}

            monthly_gb_data_scanned {
                value = 100
            }
		}
	`
}

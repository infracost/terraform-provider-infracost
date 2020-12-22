package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsVPCEndpoint(t *testing.T) {
	name := "data.infracost_aws_vpc_endpoint.vpc_endpoint"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsVPCEndpointConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_gb_data_processed", 1000),
				),
			},
		},
	})
}

func testAwsVPCEndpointConfig() string {
	return `
		data "infracost_aws_vpc_endpoint" "vpc_endpoint" {
			resources = list("vpc_endpoint_1", "vpc_endpoint_2")

			monthly_gb_data_processed {
				value = 1000
			}
		}
	`
}

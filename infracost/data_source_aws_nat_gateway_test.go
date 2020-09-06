package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsNatGateway(t *testing.T) {
	name := "data.infracost_aws_nat_gateway.my_nat_gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsNatGatewayConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_gb_data_processed", 10.0),
				),
			},
		},
	})
}

func testAwsNatGatewayConfig() string {
	return `
		data "infracost_aws_nat_gateway" "my_nat_gateway" {
			resources = list("my_nat_gateway_1", "my_nat_gateway_2")

			monthly_gb_data_processed {
				value = 10
			}
		}
	`
}

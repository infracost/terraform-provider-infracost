package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsNatGateway(t *testing.T) {
	name := "infracost_aws_nat_gateway.my-nat-gateway"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsNatGatewayConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "gb_data_processed_monthly", 10.0),
				),
			},
		},
	})
}

func testAwsNatGatewayConfig() string {
	return `
		resource "infracost_aws_nat_gateway" "my-nat-gateway" {
			resources = list("my-nat-gateway-1", "my-nat-gateway-2")

			gb_data_processed_monthly {
				value = 10
			}
		}
	`
}

package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsDXGatewayAssociationQueue(t *testing.T) {
	name := "data.infracost_aws_dx_gateway_association.dx_tgw_association"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsDXGatewayAssociationConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_gb_data_processed", 1000),
				),
			},
		},
	})
}

func testAwsDXGatewayAssociationConfig() string {
	return `
		data "infracost_aws_dx_gateway_association" "dx_tgw_association" {
			resources = list("dx_tgw_association_1", "dx_tgw_association_2")

			monthly_gb_data_processed {
				value = 1000
			}
		}
	`
}

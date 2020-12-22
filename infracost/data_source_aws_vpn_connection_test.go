package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsVPNConnection(t *testing.T) {
	name := "data.infracost_aws_vpn_connection.vpn_connection"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsVPNConnectionConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_gb_data_processed", 100),
				),
			},
		},
	})
}

func testAwsVPNConnectionConfig() string {
	return `
		data "infracost_aws_vpn_connection" "vpn_connection" {
			resources = list("vpn_connection_1", "vpn_connection_2")

			monthly_gb_data_processed {
				value = 100
			}
		}
	`
}

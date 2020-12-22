package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsEC2TransitGatewayVPCAttachment(t *testing.T) {
	name := "data.infracost_aws_ec2_transit_gateway_vpc_attachment.ec2_tgw_vpc_attach"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsEC2TransitGatewayVPCAttachmentConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_db_data_processed", 1000),
				),
			},
		},
	})
}

func testAwsEC2TransitGatewayVPCAttachmentConfig() string {
	return `
		data "infracost_aws_ec2_transit_gateway_vpc_attachment" "ec2_tgw_vpc_attach" {
			resources = list("ec2_tgw_vpc_attach_1", "ec2_tgw_vpc_attach_2")

			monthly_gb_data_processed {
				value = 1000
			}
		}
	`
}

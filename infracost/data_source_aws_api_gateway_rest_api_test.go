package infracost

import (
    "testing"

    "github.com/hashicorp/terraform/helper/resource"
)

func TestAwsApiGatewayRestApi(t *testing.T) {
    name := "data.infracost_aws_api_gateway_rest_api.my_api_requests"

    resource.Test(t, resource.TestCase{
        PreCheck:  func() {},
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAwsApiGatewayConfig(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr(name, "resources.#", "2"),
                    testCheckResourceAttrValue(name, "monthly_requests", 1000000),
                ),
            },
        },
    })
}

func testAwsApiGatewayConfig() string {
    return `
        data "infracost_aws_api_gateway_rest_api" "my_api_requests" {
            resources = list("my_rest_api_1", "my_rest_api_2")

            monthly_requests {
              value = 1000000
            }
        }
   `
}

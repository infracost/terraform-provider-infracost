package infracost

import (
    "testing"

    "github.com/hashicorp/terraform/helper/resource"
)

func TestAwsApiGatewayV2Api(t *testing.T) {
    http := "data.infracost_aws_apigatewayv2_api.http_requests"
    websocket := "data.infracost_aws_apigatewayv2_api.websocket_messages"

    resource.Test(t, resource.TestCase{
        PreCheck:  func() {},
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            {
                Config: testAwsApiGatewayV2Config(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr(http, "resources.#", "2"),
                    testCheckResourceAttrValue(http, "monthly_requests", 1000000),
                    testCheckResourceAttrValue(http, "request_size", 512),
                    resource.TestCheckResourceAttr(websocket, "resources.#", "2"),
                    testCheckResourceAttrValue(websocket, "monthly_messages", 1000000),
                    testCheckResourceAttrValue(websocket, "average_message_size", 32),
                ),
            },
        },
    })
}

func testAwsApiGatewayV2Config() string {
    return `
        data "infracost_aws_apigatewayv2_api" "http_requests" {
            resources = list("http_api_1", "http_api_2")

            monthly_requests {
              value = 1000000
            }

            request_size {
              value = 512
            }
        }

        data "infracost_aws_apigatewayv2_api" "websocket_messages" {
            resources = list("my_rest_api_1", "my_rest_api_2")

            monthly_messages {
              value = 1000000
            }

            average_message_size {
              value = 32
            }
        }
   `
}

package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAWSLambdaFunction(t *testing.T) {
	name := "infracost_aws_lambda_function.my_lambda_function"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAWSLambdaFunctionConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_requests", 100000.0),
					testCheckResourceAttrValue(name, "average_request_duration", 500.0),
				),
			},
		},
	})
}

func testAWSLambdaFunctionConfig() string {
	return `
		resource "infracost_aws_lambda_function" "my_lambda_function" {
			resources = list("my_lambda_function_1", "my_lambda_function_2")

			monthly_requests {
				value = 100000
			}

			average_request_duration {
				value = 500
			}
		}
	`
}

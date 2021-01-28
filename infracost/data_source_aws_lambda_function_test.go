package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAWSLambdaFunction(t *testing.T) {
	name := "data.infracost_aws_lambda_function.my_lambda_function"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAWSLambdaFunctionConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_requests", 100000.0),
					testCheckResourceAttrValue(name, "request_duration_ms", 500.0),
				),
			},
		},
	})
}

func testAWSLambdaFunctionConfig() string {
	return `
		data "infracost_aws_lambda_function" "my_lambda_function" {
			resources = list("my_lambda_function_1", "my_lambda_function_2")

			monthly_requests {
				value = 100000
			}

			request_duration_ms {
				value = 500
			}
		}
	`
}

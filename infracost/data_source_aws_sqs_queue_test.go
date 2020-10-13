package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsSQSQueue(t *testing.T) {
	name := "data.infracost_aws_sqs_queue.my_sqs_queue"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsSQSQueueConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_requests", 1000000),
					testCheckResourceAttrValue(name, "request_size", 64),
				),
			},
		},
	})
}

func testAwsSQSQueueConfig() string {
	return `
		data "infracost_aws_sqs_queue" "my_sqs_queue" {
			resources = list("my_sqs_queue_1", "my_sqs_queue_2")

			monthly_requests {
				value = 1000000
			}
			
			request_size {
				value = 64
			}
		}
	`
}

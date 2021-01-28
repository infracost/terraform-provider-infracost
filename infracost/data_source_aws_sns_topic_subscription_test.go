package infracost

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAwsSNSTopicSubscription(t *testing.T) {
	name := "data.infracost_aws_sns_topic_subscription.sns_subscription"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAwsSNSTopicSubscriptionConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "resources.#", "2"),
					testCheckResourceAttrValue(name, "monthly_requests", 1000000),
					testCheckResourceAttrValue(name, "request_size_kb", 64),
				),
			},
		},
	})
}

func testAwsSNSTopicSubscriptionConfig() string {
	return `
		data "infracost_aws_sns_topic_subscription" "sns_subscription" {
			resources = list("sns_topic_subscription_1", "sns_topic_subscription_2")

			monthly_requests {
				value = 1000000
			}

			request_size_kb {
				value = 64
			}
		}
	`
}

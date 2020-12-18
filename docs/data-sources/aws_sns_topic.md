# infracost_aws_sns_topic

Provides estimated usage data for an AWS SNS Topic Requests.

## Example Usage

```hcl
resource "aws_sns_topic" "sns_topic" {
  name = "sns-topic"
}

data "infracost_aws_sns_queue" "costs" {
  resources = list(aws_sns_topic.topic.id,)

  monthly_requests {
    value = 1000000
  }
  
  request_size {
    value = 64
  }

}
```

## Argument Reference

* `resources` - (Required) The IDs of the SNS Queues to apply the estimated usage.
* `monthly_requests` - (Optional) The estimated monthly requests to SNS.
* `request_size` - (Optional) The size of the requests to SNS, SNS bills in 64KB chunks. So if you process 1,000,000 requests at 128KB you pay for 2,000,000 requests.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.


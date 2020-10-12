# infracost_aws_sqs_queue

Provides estimated usage data for an AWS SQS Queue Reuqests.

## Example Usage

```hcl
resource "aws_sqs_queue" "standard" {
  name = "standard-queue"
  fifo_queue = false
}

resource "aws_sqs_queue" "fifo" {
  name = "fifo-queue"
  fifo_queue = true
}

data "infracost_aws_sqs_queue" "sqs_costs" {
  resources = list(aws_sqs_queue.standard.id, aws_sqs_qeue.fifo.id)

  monthly_requests {
    value = 1000000
  }
  
  request_size {
    value = 64
  }

}
```

## Argument Reference

* `resources` - (Required) The IDs of the SQS Queues to apply the estimated usage.
* `monthly_requests` - (Optional) The estimate in number of requests to SQS
* `request_size` - (Optional) The size of the requests to SQS, SQS bills in 64kb chunks. So if you process 1000000 requests at 128kb you pay for 2000000 requests.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.


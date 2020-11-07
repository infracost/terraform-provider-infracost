# infracost_aws_api_gateway_rest_api

Provides estimated usage data for an AWS API Gateway Rest API Requests.

## Example Usage

```hcl
resource "aws_api_gateway_rest_api" "api" {
  name = "my-rest-api"
}

data "infracost_aws_api_gateway_rest_api" "api_requests" {
  resources = list(aws_api_gateway_rest_api.api.id)

  monthly_requests {
    value = 100000000
  }
}
```

## Argument Reference

* `resources` - (Required) The ID of the Rest API.
* `monthly_requests` - (Optional) The estimated monthly requests to the Rest API Gateway.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.


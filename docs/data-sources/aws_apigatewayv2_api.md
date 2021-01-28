# infracost_aws_apigatewayv2_api

Provides estimated usage data for an AWS API Gateway v2 API. This resource supports the HTTP & Websockets protocol.

## Example Usage

```hcl
resource "aws_apigatewayv2_api" "http" {
  name = "my-http-api"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_api" "websocket" {
  name = "my-websocket-api"
  protocol_type = "WEBSOCKET"
}

data "infracost_aws_api_gateway_rest_api" "http_requests" {
  resources = list(aws_apigatewayv2_api.http.id)

  monthly_requests {
    value = 100000000
  }

  request_size_kb {
    value = 512
  }
}

data "infracost_aws_api_gateway_rest_api" "websocket_messages" {
  resources = list(aws_apigatewayv2_api.websocket.id)

  monthly_messages {
    value = 100000000
  }

  message_size_kb {
    value = 32
  }
}

```

## Argument Reference

* `resources` - (Required) The ID of the Rest API.
* `monthly_requests` - (Optional) The estimated monthly requests to the HTTP API Gateway.
* `request_size_kb` - (Optional) The estimated average request size sent to the HTTP API Gateway in KB. Requests are metered in 512KB increments, maximum size is 10MB.
* `monthly_messages` - (Optional) The number of messages sent to the Websocket API Gateway.
* `message_size_kb` - (Optional) The average size of the messages sent to the Websocket API Gateway in KB. Messages are metered in 32 KB increments, maximum size is 128KB.

### Usage values

Each of the usage value blocks currently supports the following attributes:
* `value` - (Optional) The estimated value.

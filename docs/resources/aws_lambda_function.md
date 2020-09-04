# infracost_aws_lambda_function

Provides estimated usage data for an AWS Lambda function.

## Example Usage

```hcl
resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "my_lambda_function" {
  filename      = "lambda_function_payload.zip"
  function_name = "lambda_function_name"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "exports.test"

  source_code_hash = filebase64sha256("lambda_function_payload.zip")

  runtime = "nodejs12.x"

  environment {
    variables = {
      foo = "bar"
    }
  }
}

resource "infracost_aws_lambda_function" "my_lambda_function" {
  resources = list(aws_lambda_function.my_lambda_function.id)

  monthly_requests {
    value = 100000
  }

  request_duration {
    value = 500
  }
}
```

## Argument Reference

* `resources` - (Required) The IDs of the Lambda functions to apply the estimated usage.
* `monthly_requests` - (Optional) The estimated monthly requests to the Lambda function per month. See [Usage values](#usage_values) below for details on attributes.
* `request_duration` - (Optional) The estimated duration of each request in ms. See [Usage values](#usage_values) below for details on attributes.

### Usage values

Each of the usage value blocks currently supports the following attributes:

* `value` - (Optional) The estimated value.

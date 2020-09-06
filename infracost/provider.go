package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"infracost_aws_nat_gateway":     resourceAwsNatGateway(),
			"infracost_aws_lambda_function": resourceAwsLambdaFunction(),
			"infracost_aws_dynamodb_table":  resourceAwsDynamoDB(),
		},
	}
}

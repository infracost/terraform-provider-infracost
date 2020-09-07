package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"infracost_aws_nat_gateway":     dataSourceAwsNatGateway(),
			"infracost_aws_lambda_function": dataSourceAwsLambdaFunction(),
			"infracost_aws_dynamodb_table":  dataSourceAwsDynamoDBTable(),
		},
	}
}

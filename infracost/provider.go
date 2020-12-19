package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"infracost_aws_api_gateway_rest_api": dataSourceAwsApiGatewayRestApi(),
			"infracost_aws_apigatewayv2_api":     dataSourceAwsApiGatewayV2Api(),
			"infracost_aws_codebuild_project":    dataSourceAwsCodebuildProject(),
			"infracost_aws_dynamodb_table":       dataSourceAwsDynamoDBTable(),
			"infracost_aws_cloudwatch_log_group": dataSourceAwsCloudwatchLogGroup(),
			"infracost_aws_lambda_function":      dataSourceAwsLambdaFunction(),
			"infracost_aws_nat_gateway":          dataSourceAwsNatGateway(),
			"infracost_aws_sqs_queue":            dataSourceAwsSQSQueue(),
			"infracost_aws_vpn_connection":       dataSourceAwsVPNConnection(),
		},
	}
}

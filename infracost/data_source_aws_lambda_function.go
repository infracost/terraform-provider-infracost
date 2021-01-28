package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsLambdaFunction() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":           resourcesSchema(),
			"monthly_requests":    usageSchema(),
			"request_duration_ms": usageSchema(),
		},
	}
}

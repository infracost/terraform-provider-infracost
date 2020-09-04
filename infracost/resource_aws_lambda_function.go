package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsLambdaFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"resources":                resourcesSchema(),
			"monthly_requests":         usageSchema(),
			"average_request_duration": usageSchema(),
		},
	}
}

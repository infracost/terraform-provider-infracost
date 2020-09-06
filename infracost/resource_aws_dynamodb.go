package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsDynamoDB() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"resources":                           resourcesSchema(),
			"monthly_million_write_request_units": usageSchema(),
			"monthly_million_read_request_units":  usageSchema(),
		},
	}
}

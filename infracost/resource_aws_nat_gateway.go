package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsNatGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"resources":                 resourcesSchema(),
			"monthly_gb_data_processed": usageSchema(),
		},
	}
}

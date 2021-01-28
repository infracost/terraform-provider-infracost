package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsNatGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":                 resourcesSchema(),
			"monthly_data_processed_gb": usageSchema(),
		},
	}
}

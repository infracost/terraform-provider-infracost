package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsDXGatewayAssociation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":        resourcesSchema(),
			"monthly_gb_data_processed": usageSchema(),
		},
	}
}

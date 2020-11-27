package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsCloudwatchLogGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":                 resourcesSchema(),
			"monthly_gb_data_ingestion": usageSchema(),
			"monthly_gb_data_storage":   usageSchema(),
			"monthly_gb_data_scanned":   usageSchema(),
		},
	}
}

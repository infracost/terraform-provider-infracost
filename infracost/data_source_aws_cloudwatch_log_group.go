package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsCloudwatchLogGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":                resourcesSchema(),
			"monthly_data_ingested_gb": usageSchema(),
			"storage_gb":               usageSchema(),
			"monthly_data_scanned_gb":  usageSchema(),
		},
	}
}

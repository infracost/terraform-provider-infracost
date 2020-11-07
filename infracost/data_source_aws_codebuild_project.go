package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsCodebuildProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":             resourcesSchema(),
			"monthly_build_minutes": usageSchema(),
		},
	}
}

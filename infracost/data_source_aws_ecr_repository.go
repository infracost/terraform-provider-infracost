package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsECRRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":        resourcesSchema(),
			"storage_size":     usageSchema(),
		},
	}
}

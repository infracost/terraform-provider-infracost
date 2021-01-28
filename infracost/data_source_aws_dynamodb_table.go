package infracost

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceAwsDynamoDBTable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"resources":                          resourcesSchema(),
			"monthly_write_request_units":        usageSchema(),
			"monthly_read_request_units":         usageSchema(),
			"storage_gb":                         usageSchema(),
			"pitr_backup_storage_gb":             usageSchema(),
			"on_demand_backup_storage_gb":        usageSchema(),
			"monthly_data_restored_gb":           usageSchema(),
			"monthly_streams_read_request_units": usageSchema(),
		},
	}
}

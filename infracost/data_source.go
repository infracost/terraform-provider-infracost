package infracost

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcesSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Required: true,
		Set:      schema.HashString,
	}
}

func usageSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"value": {
					Type:     schema.TypeFloat,
					Optional: true,
				},
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {
	d.SetId(uuid.New().String())
	return nil
}

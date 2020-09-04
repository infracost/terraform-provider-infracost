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

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(uuid.New().String())
	return nil
}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

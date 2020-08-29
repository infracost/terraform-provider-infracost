package infracost

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
)

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

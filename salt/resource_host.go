package salt

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostCreate,
		Read:   schema.Noop,
		Update: schema.Noop,
		Delete: schema.RemoveFromState,

		// See https://docs.saltstack.com/en/latest/topics/ssh/roster.html
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"passwd": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// Optional parameters
		},
	}
}

func resourceHostCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("host").(string))

	return nil
}

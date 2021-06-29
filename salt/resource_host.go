package salt

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostCreate,
		Read:   schema.Noop,
		Update: schema.Noop,
		Delete: schema.RemoveFromState,

		// See https://docs.saltstack.com/en/latest/topics/ssh/roster.html
		Schema: map[string]*schema.Schema{
			"salt_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"host": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Is this really required?
			"passwd": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional parameters
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sudo": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sudo_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tty": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"priv": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minion_opts": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"thin_dir": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cmd_umask": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceHostCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(d.Get("salt_id").(string))

	return nil
}

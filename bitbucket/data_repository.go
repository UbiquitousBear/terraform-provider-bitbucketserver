package bitbucket

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"forkable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"public": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fork_repository_project": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fork_repository_slug": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"enable_git_lfs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"clone_ssh": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clone_https": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRepositoryRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	project := d.Get("project").(string)
	d.SetId(fmt.Sprintf("%s/%s", project, name))
	return resourceRepositoryRead(d, m)
}

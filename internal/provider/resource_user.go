package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a Kestra User.",

		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"username": {
				Description: "The user name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The user description.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"first_name": {
				Description: "The user first name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"last_name": {
				Description: "The user last name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"email": {
				Description: "The user email.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"groups": {
				Description: "The user groups id.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	body, err := userSchemaToApi(d)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := c.request("POST", "/api/v1/users", body)
	if err != nil {
		return diag.FromErr(err)
	}

	errs := userApiToSchema(r.(map[string]interface{}), d)
	if errs != nil {
		return errs
	}

	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	userId := d.Id()

	r, err := c.request("GET", fmt.Sprintf("/api/v1/users/%s", userId), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	errs := userApiToSchema(r.(map[string]interface{}), d)
	if errs != nil {
		return errs
	}

	return diags
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	if d.HasChanges("username", "description", "first_name", "last_name", "email", "groups") {
		body, err := userSchemaToApi(d)
		if err != nil {
			return diag.FromErr(err)
		}

		userId := d.Id()

		r, err := c.request("PUT", fmt.Sprintf("/api/v1/users/%s", userId), body)
		if err != nil {
			return diag.FromErr(err)
		}

		errs := userApiToSchema(r.(map[string]interface{}), d)
		if errs != nil {
			return errs
		}

		return diags
	} else {
		return resourceUserRead(ctx, d, meta)
	}
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*Client)
	var diags diag.Diagnostics

	userId := d.Id()

	_, err := c.request("DELETE", fmt.Sprintf("/api/v1/users/%s", userId), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r TeamsLocationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Required: true,
					},
					"name": schema.StringAttribute{
						Description: "The name of the location.",
						Required:    true,
					},
					"client_default": schema.BoolAttribute{
						Description: "True if the location is the default location.",
						Optional:    true,
					},
					"ecs_support": schema.BoolAttribute{
						Description: "True if the location needs to resolve EDNS queries.",
						Optional:    true,
					},
					"networks": schema.ListNestedAttribute{
						Description: "A list of network ranges that requests from this location would originate from.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"network": schema.StringAttribute{
									Description: "The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.",
									Required:    true,
								},
							},
						},
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"doh_subdomain": schema.StringAttribute{
						Description: "The DNS over HTTPS domain to send DNS requests to. This field is auto-generated by Gateway.",
						Computed:    true,
					},
					"ip": schema.StringAttribute{
						Description: "IPV6 destination ip assigned to this location. DNS requests sent to this IP will counted as the request under this location. This field is auto-generated by Gateway.",
						Computed:    true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state TeamsLocationModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
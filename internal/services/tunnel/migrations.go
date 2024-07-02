// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r TunnelResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "UUID of the tunnel.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Cloudflare account ID",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"name": schema.StringAttribute{
						Description: "A user-friendly name for a tunnel.",
						Required:    true,
					},
					"tunnel_secret": schema.StringAttribute{
						Description: "Sets the password required to run a locally-managed tunnel. Must be at least 32 bytes and encoded as a base64 string.",
						Required:    true,
					},
					"connections": schema.ListNestedAttribute{
						Description: "The tunnel connections between your origin and Cloudflare's edge.",
						Computed:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"colo_name": schema.StringAttribute{
									Description: "The Cloudflare data center used for this connection.",
									Optional:    true,
								},
								"is_pending_reconnect": schema.BoolAttribute{
									Description: "Cloudflare continues to track connections for several minutes after they disconnect. This is an optimization to improve latency and reliability of reconnecting.  If `true`, the connection has disconnected but is still being tracked. If `false`, the connection is actively serving traffic.",
									Optional:    true,
								},
								"uuid": schema.StringAttribute{
									Description: "UUID of the Cloudflare Tunnel connection.",
									Computed:    true,
								},
							},
						},
					},
					"created_at": schema.StringAttribute{
						Description: "Timestamp of when the resource was created.",
						Computed:    true,
					},
					"deleted_at": schema.StringAttribute{
						Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state TunnelModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}

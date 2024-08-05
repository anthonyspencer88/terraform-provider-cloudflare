// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_dex_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithUpgradeState = &DeviceDEXTestResource{}

func (r *DeviceDEXTestResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "The name of the DEX test. Must be unique.",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"name": schema.StringAttribute{
						Description:   "The name of the DEX test. Must be unique.",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"account_id": schema.StringAttribute{
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"enabled": schema.BoolAttribute{
						Description: "Determines whether or not the test is active.",
						Required:    true,
					},
					"interval": schema.StringAttribute{
						Description: "How often the test will run.",
						Required:    true,
					},
					"data": schema.SingleNestedAttribute{
						Description: "The configuration object which contains the details for the WARP client to conduct the test.",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description: "The desired endpoint to test.",
								Optional:    true,
							},
							"kind": schema.StringAttribute{
								Description: "The type of test.",
								Optional:    true,
							},
							"method": schema.StringAttribute{
								Description: "The HTTP request method type.",
								Optional:    true,
							},
						},
					},
					"description": schema.StringAttribute{
						Description: "Additional details about the test.",
						Optional:    true,
					},
					"targeted": schema.BoolAttribute{
						Optional: true,
					},
					"target_policies": schema.ListNestedAttribute{
						Description: "Device settings profiles targeted by this test",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The id of the device settings profile",
									Optional:    true,
								},
								"default": schema.BoolAttribute{
									Description: "Whether the profile is the account default",
									Optional:    true,
								},
								"name": schema.StringAttribute{
									Description: "The name of the device settings profile",
									Optional:    true,
								},
							},
						},
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state DeviceDEXTestModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}

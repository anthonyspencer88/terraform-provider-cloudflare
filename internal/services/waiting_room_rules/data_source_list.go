// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_rules

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type WaitingRoomRulesListDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &WaitingRoomRulesListDataSource{}

func NewWaitingRoomRulesListDataSource() datasource.DataSource {
	return &WaitingRoomRulesListDataSource{}
}

func (d *WaitingRoomRulesListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_waiting_room_rules_list"
}

func (r *WaitingRoomRulesListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WaitingRoomRulesListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WaitingRoomRulesListDataSource

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_short_lived_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &ZeroTrustAccessShortLivedCertificateDataSource{}

func (d *ZeroTrustAccessShortLivedCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"app_id": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"aud": schema.StringAttribute{
				Description: "The Application Audience (AUD) tag. Identifies the application associated with the CA.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The ID of the CA.",
				Optional:    true,
			},
			"public_key": schema.StringAttribute{
				Description: "The public key to add to your SSH server configuration.",
				Optional:    true,
			},
		},
	}
}

func (d *ZeroTrustAccessShortLivedCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
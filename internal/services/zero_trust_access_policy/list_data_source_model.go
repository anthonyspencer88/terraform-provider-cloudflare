// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessPoliciesResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessPoliciesResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessPoliciesDataSourceModel struct {
	AppID     types.String                                     `tfsdk:"app_id" path:"app_id"`
	AccountID types.String                                     `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                     `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                      `tfsdk:"max_items"`
	Result    *[]*ZeroTrustAccessPoliciesResultDataSourceModel `tfsdk:"result"`
}

func (m *ZeroTrustAccessPoliciesDataSourceModel) toListParams() (params zero_trust.AccessApplicationPolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationPolicyListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessPoliciesResultDataSourceModel struct {
	ID                           types.String                                             `tfsdk:"id" json:"id"`
	ApprovalGroups               *[]*ZeroTrustAccessPoliciesApprovalGroupsDataSourceModel `tfsdk:"approval_groups" json:"approval_groups"`
	ApprovalRequired             types.Bool                                               `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                                        `tfsdk:"created_at" json:"created_at,computed"`
	Decision                     types.String                                             `tfsdk:"decision" json:"decision"`
	Exclude                      *[]*ZeroTrustAccessPoliciesExcludeDataSourceModel        `tfsdk:"exclude" json:"exclude"`
	Include                      *[]*ZeroTrustAccessPoliciesIncludeDataSourceModel        `tfsdk:"include" json:"include"`
	IsolationRequired            types.Bool                                               `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                             `tfsdk:"name" json:"name"`
	PurposeJustificationPrompt   types.String                                             `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	PurposeJustificationRequired types.Bool                                               `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Require                      *[]*ZeroTrustAccessPoliciesRequireDataSourceModel        `tfsdk:"require" json:"require"`
	SessionDuration              types.String                                             `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                                        `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustAccessPoliciesApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type ZeroTrustAccessPoliciesExcludeDataSourceModel struct {
	Email                *ZeroTrustAccessPoliciesExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPoliciesExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPoliciesExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                             `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPoliciesExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPoliciesExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                             `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPoliciesExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPoliciesExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPoliciesExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPoliciesExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPoliciesExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPoliciesExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPoliciesExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                             `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPoliciesExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPoliciesExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPoliciesExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPoliciesExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessPoliciesExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessPoliciesExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPoliciesIncludeDataSourceModel struct {
	Email                *ZeroTrustAccessPoliciesIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPoliciesIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPoliciesIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                             `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPoliciesIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPoliciesIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                             `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPoliciesIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPoliciesIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPoliciesIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPoliciesIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPoliciesIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPoliciesIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPoliciesIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                             `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPoliciesIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPoliciesIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPoliciesIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPoliciesIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessPoliciesIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessPoliciesIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPoliciesRequireDataSourceModel struct {
	Email                *ZeroTrustAccessPoliciesRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPoliciesRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPoliciesRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                             `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPoliciesRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPoliciesRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                             `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPoliciesRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPoliciesRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPoliciesRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPoliciesRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPoliciesRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPoliciesRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPoliciesRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                             `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPoliciesRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPoliciesRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPoliciesRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPoliciesRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessPoliciesRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessPoliciesRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessPoliciesRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
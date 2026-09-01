package awsmanagedgrafana

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGrafanaWorkspaceSamlConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#editor_role_values AwsGrafanaWorkspaceSamlConfiguration#editor_role_values}.
	// Experimental.
	EditorRoleValues *[]*string `field:"required" json:"editorRoleValues" yaml:"editorRoleValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#workspace_id AwsGrafanaWorkspaceSamlConfiguration#workspace_id}.
	// Experimental.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#admin_role_values AwsGrafanaWorkspaceSamlConfiguration#admin_role_values}.
	// Experimental.
	AdminRoleValues *[]*string `field:"optional" json:"adminRoleValues" yaml:"adminRoleValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#allowed_organizations AwsGrafanaWorkspaceSamlConfiguration#allowed_organizations}.
	// Experimental.
	AllowedOrganizations *[]*string `field:"optional" json:"allowedOrganizations" yaml:"allowedOrganizations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#email_assertion AwsGrafanaWorkspaceSamlConfiguration#email_assertion}.
	// Experimental.
	EmailAssertion *string `field:"optional" json:"emailAssertion" yaml:"emailAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#groups_assertion AwsGrafanaWorkspaceSamlConfiguration#groups_assertion}.
	// Experimental.
	GroupsAssertion *string `field:"optional" json:"groupsAssertion" yaml:"groupsAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#id AwsGrafanaWorkspaceSamlConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#idp_metadata_url AwsGrafanaWorkspaceSamlConfiguration#idp_metadata_url}.
	// Experimental.
	IdpMetadataUrl *string `field:"optional" json:"idpMetadataUrl" yaml:"idpMetadataUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#idp_metadata_xml AwsGrafanaWorkspaceSamlConfiguration#idp_metadata_xml}.
	// Experimental.
	IdpMetadataXml *string `field:"optional" json:"idpMetadataXml" yaml:"idpMetadataXml"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#login_assertion AwsGrafanaWorkspaceSamlConfiguration#login_assertion}.
	// Experimental.
	LoginAssertion *string `field:"optional" json:"loginAssertion" yaml:"loginAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#login_validity_duration AwsGrafanaWorkspaceSamlConfiguration#login_validity_duration}.
	// Experimental.
	LoginValidityDuration *float64 `field:"optional" json:"loginValidityDuration" yaml:"loginValidityDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#name_assertion AwsGrafanaWorkspaceSamlConfiguration#name_assertion}.
	// Experimental.
	NameAssertion *string `field:"optional" json:"nameAssertion" yaml:"nameAssertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#org_assertion AwsGrafanaWorkspaceSamlConfiguration#org_assertion}.
	// Experimental.
	OrgAssertion *string `field:"optional" json:"orgAssertion" yaml:"orgAssertion"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#region AwsGrafanaWorkspaceSamlConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#role_assertion AwsGrafanaWorkspaceSamlConfiguration#role_assertion}.
	// Experimental.
	RoleAssertion *string `field:"optional" json:"roleAssertion" yaml:"roleAssertion"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/grafana_workspace_saml_configuration#timeouts AwsGrafanaWorkspaceSamlConfiguration#timeouts}
	// Experimental.
	Timeouts *AwsGrafanaWorkspaceSamlConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}


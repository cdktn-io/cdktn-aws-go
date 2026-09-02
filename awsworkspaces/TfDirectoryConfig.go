package awsworkspaces

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDirectoryConfig struct {
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
	// active_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#active_directory_config TfDirectory#active_directory_config}
	// Experimental.
	ActiveDirectoryConfig *TfDirectory_ActiveDirectoryConfigProperty `field:"optional" json:"activeDirectoryConfig" yaml:"activeDirectoryConfig"`
	// certificate_based_auth_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#certificate_based_auth_properties TfDirectory#certificate_based_auth_properties}
	// Experimental.
	CertificateBasedAuthProperties *TfDirectory_CertificateBasedAuthPropertiesProperty `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#directory_id TfDirectory#directory_id}.
	// Experimental.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#id TfDirectory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#ip_group_ids TfDirectory#ip_group_ids}.
	// Experimental.
	IpGroupIds *[]*string `field:"optional" json:"ipGroupIds" yaml:"ipGroupIds"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#region TfDirectory#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// saml_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#saml_properties TfDirectory#saml_properties}
	// Experimental.
	SamlProperties *TfDirectory_SamlPropertiesProperty `field:"optional" json:"samlProperties" yaml:"samlProperties"`
	// self_service_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#self_service_permissions TfDirectory#self_service_permissions}
	// Experimental.
	SelfServicePermissions *TfDirectory_SelfServicePermissionsProperty `field:"optional" json:"selfServicePermissions" yaml:"selfServicePermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#subnet_ids TfDirectory#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#tags TfDirectory#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#tags_all TfDirectory#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#tenancy TfDirectory#tenancy}.
	// Experimental.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#user_identity_type TfDirectory#user_identity_type}.
	// Experimental.
	UserIdentityType *string `field:"optional" json:"userIdentityType" yaml:"userIdentityType"`
	// workspace_access_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#workspace_access_properties TfDirectory#workspace_access_properties}
	// Experimental.
	WorkspaceAccessProperties *TfDirectory_WorkspaceAccessPropertiesProperty `field:"optional" json:"workspaceAccessProperties" yaml:"workspaceAccessProperties"`
	// workspace_creation_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#workspace_creation_properties TfDirectory#workspace_creation_properties}
	// Experimental.
	WorkspaceCreationProperties *TfDirectory_WorkspaceCreationPropertiesProperty `field:"optional" json:"workspaceCreationProperties" yaml:"workspaceCreationProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#workspace_directory_description TfDirectory#workspace_directory_description}.
	// Experimental.
	WorkspaceDirectoryDescription *string `field:"optional" json:"workspaceDirectoryDescription" yaml:"workspaceDirectoryDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#workspace_directory_name TfDirectory#workspace_directory_name}.
	// Experimental.
	WorkspaceDirectoryName *string `field:"optional" json:"workspaceDirectoryName" yaml:"workspaceDirectoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#workspace_type TfDirectory#workspace_type}.
	// Experimental.
	WorkspaceType *string `field:"optional" json:"workspaceType" yaml:"workspaceType"`
}


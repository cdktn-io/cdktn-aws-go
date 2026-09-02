package awsworkspacesweb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserAccessLoggingSettingsAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_access_logging_settings_association#portal_arn TfUserAccessLoggingSettingsAssociation#portal_arn}.
	// Experimental.
	PortalArn *string `field:"required" json:"portalArn" yaml:"portalArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_access_logging_settings_association#user_access_logging_settings_arn TfUserAccessLoggingSettingsAssociation#user_access_logging_settings_arn}.
	// Experimental.
	UserAccessLoggingSettingsArn *string `field:"required" json:"userAccessLoggingSettingsArn" yaml:"userAccessLoggingSettingsArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_user_access_logging_settings_association#region TfUserAccessLoggingSettingsAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}


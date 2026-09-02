package awsworkspacesweb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNetworkSettingsAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_network_settings_association#network_settings_arn TfNetworkSettingsAssociation#network_settings_arn}.
	// Experimental.
	NetworkSettingsArn *string `field:"required" json:"networkSettingsArn" yaml:"networkSettingsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_network_settings_association#portal_arn TfNetworkSettingsAssociation#portal_arn}.
	// Experimental.
	PortalArn *string `field:"required" json:"portalArn" yaml:"portalArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_network_settings_association#region TfNetworkSettingsAssociation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}


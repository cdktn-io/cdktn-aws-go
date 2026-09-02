package awskendra

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIndexConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#name TfIndex#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#role_arn TfIndex#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// capacity_units block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#capacity_units TfIndex#capacity_units}
	// Experimental.
	CapacityUnits *TfIndex_CapacityUnitsProperty `field:"optional" json:"capacityUnits" yaml:"capacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#description TfIndex#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// document_metadata_configuration_updates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#document_metadata_configuration_updates TfIndex#document_metadata_configuration_updates}
	// Experimental.
	DocumentMetadataConfigurationUpdates interface{} `field:"optional" json:"documentMetadataConfigurationUpdates" yaml:"documentMetadataConfigurationUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#edition TfIndex#edition}.
	// Experimental.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#id TfIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#region TfIndex#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// server_side_encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#server_side_encryption_configuration TfIndex#server_side_encryption_configuration}
	// Experimental.
	ServerSideEncryptionConfiguration *TfIndex_ServerSideEncryptionConfigurationProperty `field:"optional" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#tags TfIndex#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#tags_all TfIndex#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#timeouts TfIndex#timeouts}
	// Experimental.
	Timeouts *TfIndex_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#user_context_policy TfIndex#user_context_policy}.
	// Experimental.
	UserContextPolicy *string `field:"optional" json:"userContextPolicy" yaml:"userContextPolicy"`
	// user_group_resolution_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#user_group_resolution_configuration TfIndex#user_group_resolution_configuration}
	// Experimental.
	UserGroupResolutionConfiguration *TfIndex_UserGroupResolutionConfigurationProperty `field:"optional" json:"userGroupResolutionConfiguration" yaml:"userGroupResolutionConfiguration"`
	// user_token_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#user_token_configurations TfIndex#user_token_configurations}
	// Experimental.
	UserTokenConfigurations *TfIndex_UserTokenConfigurationsProperty `field:"optional" json:"userTokenConfigurations" yaml:"userTokenConfigurations"`
}


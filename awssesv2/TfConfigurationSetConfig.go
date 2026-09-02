package awssesv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigurationSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#configuration_set_name TfConfigurationSet#configuration_set_name}.
	// Experimental.
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
	// delivery_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#delivery_options TfConfigurationSet#delivery_options}
	// Experimental.
	DeliveryOptions *TfConfigurationSet_DeliveryOptionsProperty `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#id TfConfigurationSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#region TfConfigurationSet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// reputation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#reputation_options TfConfigurationSet#reputation_options}
	// Experimental.
	ReputationOptions *TfConfigurationSet_ReputationOptionsProperty `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// sending_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#sending_options TfConfigurationSet#sending_options}
	// Experimental.
	SendingOptions *TfConfigurationSet_SendingOptionsProperty `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// suppression_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#suppression_options TfConfigurationSet#suppression_options}
	// Experimental.
	SuppressionOptions *TfConfigurationSet_SuppressionOptionsProperty `field:"optional" json:"suppressionOptions" yaml:"suppressionOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#tags TfConfigurationSet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#tags_all TfConfigurationSet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tracking_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#tracking_options TfConfigurationSet#tracking_options}
	// Experimental.
	TrackingOptions *TfConfigurationSet_TrackingOptionsProperty `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
	// vdm_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#vdm_options TfConfigurationSet#vdm_options}
	// Experimental.
	VdmOptions *TfConfigurationSet_VdmOptionsProperty `field:"optional" json:"vdmOptions" yaml:"vdmOptions"`
}


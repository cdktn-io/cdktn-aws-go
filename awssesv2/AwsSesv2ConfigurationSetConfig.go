package awssesv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSesv2ConfigurationSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#configuration_set_name AwsSesv2ConfigurationSet#configuration_set_name}.
	// Experimental.
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
	// delivery_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#delivery_options AwsSesv2ConfigurationSet#delivery_options}
	// Experimental.
	DeliveryOptions *AwsSesv2ConfigurationSet_DeliveryOptionsProperty `field:"optional" json:"deliveryOptions" yaml:"deliveryOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#id AwsSesv2ConfigurationSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#region AwsSesv2ConfigurationSet#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// reputation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#reputation_options AwsSesv2ConfigurationSet#reputation_options}
	// Experimental.
	ReputationOptions *AwsSesv2ConfigurationSet_ReputationOptionsProperty `field:"optional" json:"reputationOptions" yaml:"reputationOptions"`
	// sending_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#sending_options AwsSesv2ConfigurationSet#sending_options}
	// Experimental.
	SendingOptions *AwsSesv2ConfigurationSet_SendingOptionsProperty `field:"optional" json:"sendingOptions" yaml:"sendingOptions"`
	// suppression_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#suppression_options AwsSesv2ConfigurationSet#suppression_options}
	// Experimental.
	SuppressionOptions *AwsSesv2ConfigurationSet_SuppressionOptionsProperty `field:"optional" json:"suppressionOptions" yaml:"suppressionOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#tags AwsSesv2ConfigurationSet#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#tags_all AwsSesv2ConfigurationSet#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tracking_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#tracking_options AwsSesv2ConfigurationSet#tracking_options}
	// Experimental.
	TrackingOptions *AwsSesv2ConfigurationSet_TrackingOptionsProperty `field:"optional" json:"trackingOptions" yaml:"trackingOptions"`
	// vdm_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#vdm_options AwsSesv2ConfigurationSet#vdm_options}
	// Experimental.
	VdmOptions *AwsSesv2ConfigurationSet_VdmOptionsProperty `field:"optional" json:"vdmOptions" yaml:"vdmOptions"`
}


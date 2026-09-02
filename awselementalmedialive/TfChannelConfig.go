package awselementalmedialive

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#channel_class TfChannel#channel_class}.
	// Experimental.
	ChannelClass *string `field:"required" json:"channelClass" yaml:"channelClass"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destinations TfChannel#destinations}
	// Experimental.
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// encoder_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#encoder_settings TfChannel#encoder_settings}
	// Experimental.
	EncoderSettings *TfChannel_EncoderSettingsProperty `field:"required" json:"encoderSettings" yaml:"encoderSettings"`
	// input_attachments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_attachments TfChannel#input_attachments}
	// Experimental.
	InputAttachments interface{} `field:"required" json:"inputAttachments" yaml:"inputAttachments"`
	// input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_specification TfChannel#input_specification}
	// Experimental.
	InputSpecification *TfChannel_InputSpecificationProperty `field:"required" json:"inputSpecification" yaml:"inputSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#name TfChannel#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// cdi_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#cdi_input_specification TfChannel#cdi_input_specification}
	// Experimental.
	CdiInputSpecification *TfChannel_CdiInputSpecificationProperty `field:"optional" json:"cdiInputSpecification" yaml:"cdiInputSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#id TfChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#log_level TfChannel#log_level}.
	// Experimental.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// maintenance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#maintenance TfChannel#maintenance}
	// Experimental.
	Maintenance *TfChannel_MaintenanceProperty `field:"optional" json:"maintenance" yaml:"maintenance"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#region TfChannel#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#role_arn TfChannel#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#start_channel TfChannel#start_channel}.
	// Experimental.
	StartChannel interface{} `field:"optional" json:"startChannel" yaml:"startChannel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tags TfChannel#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tags_all TfChannel#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timeouts TfChannel#timeouts}
	// Experimental.
	Timeouts *TfChannel_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#vpc TfChannel#vpc}
	// Experimental.
	Vpc *TfChannel_VpcProperty `field:"optional" json:"vpc" yaml:"vpc"`
}


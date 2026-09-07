package chime

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVoiceConnectorTerminationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#calling_regions AwsVoiceConnectorTermination#calling_regions}.
	// Experimental.
	CallingRegions *[]*string `field:"required" json:"callingRegions" yaml:"callingRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#cidr_allow_list AwsVoiceConnectorTermination#cidr_allow_list}.
	// Experimental.
	CidrAllowList *[]*string `field:"required" json:"cidrAllowList" yaml:"cidrAllowList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#voice_connector_id AwsVoiceConnectorTermination#voice_connector_id}.
	// Experimental.
	VoiceConnectorId *string `field:"required" json:"voiceConnectorId" yaml:"voiceConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#cps_limit AwsVoiceConnectorTermination#cps_limit}.
	// Experimental.
	CpsLimit *float64 `field:"optional" json:"cpsLimit" yaml:"cpsLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#default_phone_number AwsVoiceConnectorTermination#default_phone_number}.
	// Experimental.
	DefaultPhoneNumber *string `field:"optional" json:"defaultPhoneNumber" yaml:"defaultPhoneNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#disabled AwsVoiceConnectorTermination#disabled}.
	// Experimental.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#id AwsVoiceConnectorTermination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination#region AwsVoiceConnectorTermination#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}


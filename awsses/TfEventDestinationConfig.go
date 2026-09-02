package awsses

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEventDestinationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#configuration_set_name TfEventDestination#configuration_set_name}.
	// Experimental.
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#matching_types TfEventDestination#matching_types}.
	// Experimental.
	MatchingTypes *[]*string `field:"required" json:"matchingTypes" yaml:"matchingTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#name TfEventDestination#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloudwatch_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#cloudwatch_destination TfEventDestination#cloudwatch_destination}
	// Experimental.
	CloudwatchDestination interface{} `field:"optional" json:"cloudwatchDestination" yaml:"cloudwatchDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#enabled TfEventDestination#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#id TfEventDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kinesis_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#kinesis_destination TfEventDestination#kinesis_destination}
	// Experimental.
	KinesisDestination *TfEventDestination_KinesisDestinationProperty `field:"optional" json:"kinesisDestination" yaml:"kinesisDestination"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#region TfEventDestination#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// sns_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_event_destination#sns_destination TfEventDestination#sns_destination}
	// Experimental.
	SnsDestination *TfEventDestination_SnsDestinationProperty `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}


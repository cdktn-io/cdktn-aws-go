package awscodecommit


// Experimental.
type TfTrigger_TriggerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codecommit_trigger#destination_arn TfTrigger#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codecommit_trigger#events TfTrigger#events}.
	// Experimental.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codecommit_trigger#name TfTrigger#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codecommit_trigger#branches TfTrigger#branches}.
	// Experimental.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codecommit_trigger#custom_data TfTrigger#custom_data}.
	// Experimental.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
}


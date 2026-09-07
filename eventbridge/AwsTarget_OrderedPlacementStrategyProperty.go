package eventbridge


// Experimental.
type AwsTarget_OrderedPlacementStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#type AwsTarget#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#field AwsTarget#field}.
	// Experimental.
	Field *string `field:"optional" json:"field" yaml:"field"`
}


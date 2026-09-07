package eventbridge


// Experimental.
type AwsTarget_PlacementConstraintProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#type AwsTarget#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#expression AwsTarget#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}


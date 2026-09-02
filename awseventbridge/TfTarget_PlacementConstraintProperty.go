package awseventbridge


// Experimental.
type TfTarget_PlacementConstraintProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#type TfTarget#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#expression TfTarget#expression}.
	// Experimental.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}


package awsautoscaling


// Experimental.
type TfGroup_TrafficSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#identifier TfGroup#identifier}.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#type TfGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


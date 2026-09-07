package autoscaling


// Experimental.
type AwsGroup_TrafficSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#identifier AwsGroup#identifier}.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#type AwsGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


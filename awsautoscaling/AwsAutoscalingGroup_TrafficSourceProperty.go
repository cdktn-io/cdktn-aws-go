package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_TrafficSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#identifier AwsAutoscalingGroup#identifier}.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#type AwsAutoscalingGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


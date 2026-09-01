package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_TotalLocalStorageGbProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#max AwsAutoscalingGroup#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#min AwsAutoscalingGroup#min}.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}


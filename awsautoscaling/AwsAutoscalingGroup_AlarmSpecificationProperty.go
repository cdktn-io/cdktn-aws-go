package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_AlarmSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#alarms AwsAutoscalingGroup#alarms}.
	// Experimental.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
}


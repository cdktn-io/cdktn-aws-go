package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_AvailabilityZoneDistributionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#capacity_distribution_strategy AwsAutoscalingGroup#capacity_distribution_strategy}.
	// Experimental.
	CapacityDistributionStrategy *string `field:"optional" json:"capacityDistributionStrategy" yaml:"capacityDistributionStrategy"`
}


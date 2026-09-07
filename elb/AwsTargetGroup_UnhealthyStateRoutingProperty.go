package elb


// Experimental.
type AwsTargetGroup_UnhealthyStateRoutingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#minimum_healthy_targets_count AwsTargetGroup#minimum_healthy_targets_count}.
	// Experimental.
	MinimumHealthyTargetsCount *float64 `field:"optional" json:"minimumHealthyTargetsCount" yaml:"minimumHealthyTargetsCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#minimum_healthy_targets_percentage AwsTargetGroup#minimum_healthy_targets_percentage}.
	// Experimental.
	MinimumHealthyTargetsPercentage *string `field:"optional" json:"minimumHealthyTargetsPercentage" yaml:"minimumHealthyTargetsPercentage"`
}


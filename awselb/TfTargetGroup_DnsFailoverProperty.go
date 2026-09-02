package awselb


// Experimental.
type TfTargetGroup_DnsFailoverProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#minimum_healthy_targets_count TfTargetGroup#minimum_healthy_targets_count}.
	// Experimental.
	MinimumHealthyTargetsCount *string `field:"optional" json:"minimumHealthyTargetsCount" yaml:"minimumHealthyTargetsCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#minimum_healthy_targets_percentage TfTargetGroup#minimum_healthy_targets_percentage}.
	// Experimental.
	MinimumHealthyTargetsPercentage *string `field:"optional" json:"minimumHealthyTargetsPercentage" yaml:"minimumHealthyTargetsPercentage"`
}


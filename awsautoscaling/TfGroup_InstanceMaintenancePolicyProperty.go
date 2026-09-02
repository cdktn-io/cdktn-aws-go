package awsautoscaling


// Experimental.
type TfGroup_InstanceMaintenancePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#max_healthy_percentage TfGroup#max_healthy_percentage}.
	// Experimental.
	MaxHealthyPercentage *float64 `field:"required" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#min_healthy_percentage TfGroup#min_healthy_percentage}.
	// Experimental.
	MinHealthyPercentage *float64 `field:"required" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
}


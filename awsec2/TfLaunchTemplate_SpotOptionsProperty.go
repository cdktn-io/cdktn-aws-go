package awsec2


// Experimental.
type TfLaunchTemplate_SpotOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#block_duration_minutes TfLaunchTemplate#block_duration_minutes}.
	// Experimental.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#instance_interruption_behavior TfLaunchTemplate#instance_interruption_behavior}.
	// Experimental.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#max_price TfLaunchTemplate#max_price}.
	// Experimental.
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#spot_instance_type TfLaunchTemplate#spot_instance_type}.
	// Experimental.
	SpotInstanceType *string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#valid_until TfLaunchTemplate#valid_until}.
	// Experimental.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}


package awsec2


// Experimental.
type TfInstance_SpotOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#instance_interruption_behavior TfInstance#instance_interruption_behavior}.
	// Experimental.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#max_price TfInstance#max_price}.
	// Experimental.
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#spot_instance_type TfInstance#spot_instance_type}.
	// Experimental.
	SpotInstanceType *string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#valid_until TfInstance#valid_until}.
	// Experimental.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}


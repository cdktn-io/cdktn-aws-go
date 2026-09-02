package awsemr


// Experimental.
type TfInstanceFleet_SpotSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#allocation_strategy TfInstanceFleet#allocation_strategy}.
	// Experimental.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#timeout_action TfInstanceFleet#timeout_action}.
	// Experimental.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#timeout_duration_minutes TfInstanceFleet#timeout_duration_minutes}.
	// Experimental.
	TimeoutDurationMinutes *float64 `field:"required" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#block_duration_minutes TfInstanceFleet#block_duration_minutes}.
	// Experimental.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
}


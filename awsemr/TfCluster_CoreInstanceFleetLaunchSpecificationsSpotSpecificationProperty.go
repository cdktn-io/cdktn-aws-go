package awsemr


// Experimental.
type TfCluster_CoreInstanceFleetLaunchSpecificationsSpotSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#allocation_strategy TfCluster#allocation_strategy}.
	// Experimental.
	AllocationStrategy *string `field:"required" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#timeout_action TfCluster#timeout_action}.
	// Experimental.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#timeout_duration_minutes TfCluster#timeout_duration_minutes}.
	// Experimental.
	TimeoutDurationMinutes *float64 `field:"required" json:"timeoutDurationMinutes" yaml:"timeoutDurationMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#block_duration_minutes TfCluster#block_duration_minutes}.
	// Experimental.
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
}


package awsopensearch


// Experimental.
type AwsOpensearchDomain_AutoTuneOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#desired_state AwsOpensearchDomain#desired_state}.
	// Experimental.
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// maintenance_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#maintenance_schedule AwsOpensearchDomain#maintenance_schedule}
	// Experimental.
	MaintenanceSchedule interface{} `field:"optional" json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#rollback_on_disable AwsOpensearchDomain#rollback_on_disable}.
	// Experimental.
	RollbackOnDisable *string `field:"optional" json:"rollbackOnDisable" yaml:"rollbackOnDisable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#use_off_peak_window AwsOpensearchDomain#use_off_peak_window}.
	// Experimental.
	UseOffPeakWindow interface{} `field:"optional" json:"useOffPeakWindow" yaml:"useOffPeakWindow"`
}


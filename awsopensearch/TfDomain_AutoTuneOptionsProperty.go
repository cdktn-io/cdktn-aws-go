package awsopensearch


// Experimental.
type TfDomain_AutoTuneOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#desired_state TfDomain#desired_state}.
	// Experimental.
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// maintenance_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#maintenance_schedule TfDomain#maintenance_schedule}
	// Experimental.
	MaintenanceSchedule interface{} `field:"optional" json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#rollback_on_disable TfDomain#rollback_on_disable}.
	// Experimental.
	RollbackOnDisable *string `field:"optional" json:"rollbackOnDisable" yaml:"rollbackOnDisable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#use_off_peak_window TfDomain#use_off_peak_window}.
	// Experimental.
	UseOffPeakWindow interface{} `field:"optional" json:"useOffPeakWindow" yaml:"useOffPeakWindow"`
}


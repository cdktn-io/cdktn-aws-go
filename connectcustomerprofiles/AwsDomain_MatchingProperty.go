package connectcustomerprofiles


// Experimental.
type AwsDomain_MatchingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled AwsDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// auto_merging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#auto_merging AwsDomain#auto_merging}
	// Experimental.
	AutoMerging *AwsDomain_AutoMergingProperty `field:"optional" json:"autoMerging" yaml:"autoMerging"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#exporting_config AwsDomain#exporting_config}
	// Experimental.
	ExportingConfig *AwsDomain_MatchingExportingConfigProperty `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// job_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#job_schedule AwsDomain#job_schedule}
	// Experimental.
	JobSchedule *AwsDomain_JobScheduleProperty `field:"optional" json:"jobSchedule" yaml:"jobSchedule"`
}


package awsconnectcustomerprofiles


// Experimental.
type TfDomain_MatchingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled TfDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// auto_merging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#auto_merging TfDomain#auto_merging}
	// Experimental.
	AutoMerging *TfDomain_AutoMergingProperty `field:"optional" json:"autoMerging" yaml:"autoMerging"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#exporting_config TfDomain#exporting_config}
	// Experimental.
	ExportingConfig *TfDomain_MatchingExportingConfigProperty `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// job_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#job_schedule TfDomain#job_schedule}
	// Experimental.
	JobSchedule *TfDomain_JobScheduleProperty `field:"optional" json:"jobSchedule" yaml:"jobSchedule"`
}


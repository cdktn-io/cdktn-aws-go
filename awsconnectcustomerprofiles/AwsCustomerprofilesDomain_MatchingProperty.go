package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_MatchingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled AwsCustomerprofilesDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// auto_merging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#auto_merging AwsCustomerprofilesDomain#auto_merging}
	// Experimental.
	AutoMerging *AwsCustomerprofilesDomain_AutoMergingProperty `field:"optional" json:"autoMerging" yaml:"autoMerging"`
	// exporting_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#exporting_config AwsCustomerprofilesDomain#exporting_config}
	// Experimental.
	ExportingConfig *AwsCustomerprofilesDomain_MatchingExportingConfigProperty `field:"optional" json:"exportingConfig" yaml:"exportingConfig"`
	// job_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#job_schedule AwsCustomerprofilesDomain#job_schedule}
	// Experimental.
	JobSchedule *AwsCustomerprofilesDomain_JobScheduleProperty `field:"optional" json:"jobSchedule" yaml:"jobSchedule"`
}


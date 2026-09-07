package s3files


// Experimental.
type AwsSynchronizationConfiguration_ExpirationDataRuleProperty struct {
	// Days after last access before data expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_synchronization_configuration#days_after_last_access AwsSynchronizationConfiguration#days_after_last_access}
	// Experimental.
	DaysAfterLastAccess *float64 `field:"required" json:"daysAfterLastAccess" yaml:"daysAfterLastAccess"`
}


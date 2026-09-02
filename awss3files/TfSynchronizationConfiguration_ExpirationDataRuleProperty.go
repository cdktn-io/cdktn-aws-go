package awss3files


// Experimental.
type TfSynchronizationConfiguration_ExpirationDataRuleProperty struct {
	// Days after last access before data expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_synchronization_configuration#days_after_last_access TfSynchronizationConfiguration#days_after_last_access}
	// Experimental.
	DaysAfterLastAccess *float64 `field:"required" json:"daysAfterLastAccess" yaml:"daysAfterLastAccess"`
}


package s3files


// Experimental.
type AwsSynchronizationConfiguration_ImportDataRuleProperty struct {
	// S3 prefix for import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_synchronization_configuration#prefix AwsSynchronizationConfiguration#prefix}
	// Experimental.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Maximum file size to import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_synchronization_configuration#size_less_than AwsSynchronizationConfiguration#size_less_than}
	// Experimental.
	SizeLessThan *float64 `field:"required" json:"sizeLessThan" yaml:"sizeLessThan"`
	// Import trigger type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_synchronization_configuration#trigger AwsSynchronizationConfiguration#trigger}
	// Experimental.
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
}


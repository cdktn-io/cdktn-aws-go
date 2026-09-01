package awssagemakerai


// Experimental.
type AwsSagemakerDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_ebs_volume_size_in_gb AwsSagemakerDomain#default_ebs_volume_size_in_gb}.
	// Experimental.
	DefaultEbsVolumeSizeInGb *float64 `field:"required" json:"defaultEbsVolumeSizeInGb" yaml:"defaultEbsVolumeSizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#maximum_ebs_volume_size_in_gb AwsSagemakerDomain#maximum_ebs_volume_size_in_gb}.
	// Experimental.
	MaximumEbsVolumeSizeInGb *float64 `field:"required" json:"maximumEbsVolumeSizeInGb" yaml:"maximumEbsVolumeSizeInGb"`
}


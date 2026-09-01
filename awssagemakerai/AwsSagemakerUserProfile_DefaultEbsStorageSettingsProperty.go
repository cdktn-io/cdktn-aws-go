package awssagemakerai


// Experimental.
type AwsSagemakerUserProfile_DefaultEbsStorageSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#default_ebs_volume_size_in_gb AwsSagemakerUserProfile#default_ebs_volume_size_in_gb}.
	// Experimental.
	DefaultEbsVolumeSizeInGb *float64 `field:"required" json:"defaultEbsVolumeSizeInGb" yaml:"defaultEbsVolumeSizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#maximum_ebs_volume_size_in_gb AwsSagemakerUserProfile#maximum_ebs_volume_size_in_gb}.
	// Experimental.
	MaximumEbsVolumeSizeInGb *float64 `field:"required" json:"maximumEbsVolumeSizeInGb" yaml:"maximumEbsVolumeSizeInGb"`
}


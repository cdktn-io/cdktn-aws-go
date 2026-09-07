package sagemakerai


// Experimental.
type AwsSpace_SpaceStorageSettingsProperty struct {
	// ebs_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#ebs_storage_settings AwsSpace#ebs_storage_settings}
	// Experimental.
	EbsStorageSettings *AwsSpace_EbsStorageSettingsProperty `field:"required" json:"ebsStorageSettings" yaml:"ebsStorageSettings"`
}


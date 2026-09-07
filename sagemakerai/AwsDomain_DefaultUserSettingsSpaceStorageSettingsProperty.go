package sagemakerai


// Experimental.
type AwsDomain_DefaultUserSettingsSpaceStorageSettingsProperty struct {
	// default_ebs_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_ebs_storage_settings AwsDomain#default_ebs_storage_settings}
	// Experimental.
	DefaultEbsStorageSettings *AwsDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}


package awssagemakerai


// Experimental.
type TfUserProfile_SpaceStorageSettingsProperty struct {
	// default_ebs_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#default_ebs_storage_settings TfUserProfile#default_ebs_storage_settings}
	// Experimental.
	DefaultEbsStorageSettings *TfUserProfile_DefaultEbsStorageSettingsProperty `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}


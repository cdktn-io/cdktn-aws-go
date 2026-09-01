package awssagemakerai


// Experimental.
type AwsSagemakerDomain_DefaultUserSettingsSpaceStorageSettingsProperty struct {
	// default_ebs_storage_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_ebs_storage_settings AwsSagemakerDomain#default_ebs_storage_settings}
	// Experimental.
	DefaultEbsStorageSettings *AwsSagemakerDomain_DefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettingsProperty `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}


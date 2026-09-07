package sagemakerai


// Experimental.
type AwsUserProfile_RSessionAppSettingsProperty struct {
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#custom_image AwsUserProfile#custom_image}
	// Experimental.
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#default_resource_spec AwsUserProfile#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *AwsUserProfile_UserSettingsRSessionAppSettingsDefaultResourceSpecProperty `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}


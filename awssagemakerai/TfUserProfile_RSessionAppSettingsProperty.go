package awssagemakerai


// Experimental.
type TfUserProfile_RSessionAppSettingsProperty struct {
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#custom_image TfUserProfile#custom_image}
	// Experimental.
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#default_resource_spec TfUserProfile#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfUserProfile_UserSettingsRSessionAppSettingsDefaultResourceSpecProperty `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}


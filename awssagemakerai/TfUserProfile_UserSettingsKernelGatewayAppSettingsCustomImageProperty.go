package awssagemakerai


// Experimental.
type TfUserProfile_UserSettingsKernelGatewayAppSettingsCustomImageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#app_image_config_name TfUserProfile#app_image_config_name}.
	// Experimental.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#image_name TfUserProfile#image_name}.
	// Experimental.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#image_version_number TfUserProfile#image_version_number}.
	// Experimental.
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}


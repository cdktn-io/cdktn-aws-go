package awssagemakerai


// Experimental.
type TfDomain_DefaultUserSettingsRSessionAppSettingsCustomImageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#app_image_config_name TfDomain#app_image_config_name}.
	// Experimental.
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#image_name TfDomain#image_name}.
	// Experimental.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#image_version_number TfDomain#image_version_number}.
	// Experimental.
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}


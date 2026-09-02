package awssagemakerai


// Experimental.
type TfDomain_RSessionAppSettingsProperty struct {
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#custom_image TfDomain#custom_image}
	// Experimental.
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_resource_spec TfDomain#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfDomain_DefaultUserSettingsRSessionAppSettingsDefaultResourceSpecProperty `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}


package awssagemakerai


// Experimental.
type TfSpace_KernelGatewayAppSettingsProperty struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#default_resource_spec TfSpace#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfSpace_SpaceSettingsKernelGatewayAppSettingsDefaultResourceSpecProperty `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#custom_image TfSpace#custom_image}
	// Experimental.
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#lifecycle_config_arns TfSpace#lifecycle_config_arns}.
	// Experimental.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}


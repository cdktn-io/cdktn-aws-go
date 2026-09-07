package sagemakerai


// Experimental.
type AwsSpace_JupyterServerAppSettingsProperty struct {
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#default_resource_spec AwsSpace#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *AwsSpace_SpaceSettingsJupyterServerAppSettingsDefaultResourceSpecProperty `field:"required" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#code_repository AwsSpace#code_repository}
	// Experimental.
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#lifecycle_config_arns AwsSpace#lifecycle_config_arns}.
	// Experimental.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}


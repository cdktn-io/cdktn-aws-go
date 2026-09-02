package awssagemakerai


// Experimental.
type TfDomain_DefaultUserSettingsJupyterServerAppSettingsProperty struct {
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#code_repository TfDomain#code_repository}
	// Experimental.
	CodeRepository interface{} `field:"optional" json:"codeRepository" yaml:"codeRepository"`
	// default_resource_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#default_resource_spec TfDomain#default_resource_spec}
	// Experimental.
	DefaultResourceSpec *TfDomain_DefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecProperty `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#lifecycle_config_arns TfDomain#lifecycle_config_arns}.
	// Experimental.
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}


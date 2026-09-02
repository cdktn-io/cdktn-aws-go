package awssagemakerai


// Experimental.
type TfDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#assumable_role_arns TfDomain#assumable_role_arns}.
	// Experimental.
	AssumableRoleArns *[]*string `field:"optional" json:"assumableRoleArns" yaml:"assumableRoleArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#execution_role_arns TfDomain#execution_role_arns}.
	// Experimental.
	ExecutionRoleArns *[]*string `field:"optional" json:"executionRoleArns" yaml:"executionRoleArns"`
}


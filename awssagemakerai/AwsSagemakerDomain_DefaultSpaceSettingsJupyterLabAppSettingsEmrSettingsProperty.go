package awssagemakerai


// Experimental.
type AwsSagemakerDomain_DefaultSpaceSettingsJupyterLabAppSettingsEmrSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#assumable_role_arns AwsSagemakerDomain#assumable_role_arns}.
	// Experimental.
	AssumableRoleArns *[]*string `field:"optional" json:"assumableRoleArns" yaml:"assumableRoleArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#execution_role_arns AwsSagemakerDomain#execution_role_arns}.
	// Experimental.
	ExecutionRoleArns *[]*string `field:"optional" json:"executionRoleArns" yaml:"executionRoleArns"`
}


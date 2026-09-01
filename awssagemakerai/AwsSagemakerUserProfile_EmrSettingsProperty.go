package awssagemakerai


// Experimental.
type AwsSagemakerUserProfile_EmrSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#assumable_role_arns AwsSagemakerUserProfile#assumable_role_arns}.
	// Experimental.
	AssumableRoleArns *[]*string `field:"optional" json:"assumableRoleArns" yaml:"assumableRoleArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_user_profile#execution_role_arns AwsSagemakerUserProfile#execution_role_arns}.
	// Experimental.
	ExecutionRoleArns *[]*string `field:"optional" json:"executionRoleArns" yaml:"executionRoleArns"`
}


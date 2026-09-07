package codepipeline


// Experimental.
type AwsCodepipeline_StageOnSuccessConditionRuleRuleTypeIdProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#category AwsCodepipeline#category}.
	// Experimental.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#provider AwsCodepipeline#provider}.
	// Experimental.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#owner AwsCodepipeline#owner}.
	// Experimental.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#version AwsCodepipeline#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}


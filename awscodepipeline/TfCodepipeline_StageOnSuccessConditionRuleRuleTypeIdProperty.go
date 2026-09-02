package awscodepipeline


// Experimental.
type TfCodepipeline_StageOnSuccessConditionRuleRuleTypeIdProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#category TfCodepipeline#category}.
	// Experimental.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#provider TfCodepipeline#provider}.
	// Experimental.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#owner TfCodepipeline#owner}.
	// Experimental.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline#version TfCodepipeline#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}


package awscodepipeline


// Experimental.
type AwsCodepipelineCustomActionType_ConfigurationPropertyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#key AwsCodepipelineCustomActionType#key}.
	// Experimental.
	Key interface{} `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#name AwsCodepipelineCustomActionType#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#required AwsCodepipelineCustomActionType#required}.
	// Experimental.
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#secret AwsCodepipelineCustomActionType#secret}.
	// Experimental.
	Secret interface{} `field:"required" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#description AwsCodepipelineCustomActionType#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#queryable AwsCodepipelineCustomActionType#queryable}.
	// Experimental.
	Queryable interface{} `field:"optional" json:"queryable" yaml:"queryable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#type AwsCodepipelineCustomActionType#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


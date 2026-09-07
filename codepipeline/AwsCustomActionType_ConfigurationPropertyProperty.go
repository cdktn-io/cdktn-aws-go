package codepipeline


// Experimental.
type AwsCustomActionType_ConfigurationPropertyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#key AwsCustomActionType#key}.
	// Experimental.
	Key interface{} `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#name AwsCustomActionType#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#required AwsCustomActionType#required}.
	// Experimental.
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#secret AwsCustomActionType#secret}.
	// Experimental.
	Secret interface{} `field:"required" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#description AwsCustomActionType#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#queryable AwsCustomActionType#queryable}.
	// Experimental.
	Queryable interface{} `field:"optional" json:"queryable" yaml:"queryable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#type AwsCustomActionType#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


package codepipeline


// Experimental.
type AwsCustomActionType_OutputArtifactDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#maximum_count AwsCustomActionType#maximum_count}.
	// Experimental.
	MaximumCount *float64 `field:"required" json:"maximumCount" yaml:"maximumCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_custom_action_type#minimum_count AwsCustomActionType#minimum_count}.
	// Experimental.
	MinimumCount *float64 `field:"required" json:"minimumCount" yaml:"minimumCount"`
}


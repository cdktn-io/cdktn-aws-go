package awschimesdkmediapipelines


// Experimental.
type TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#keywords TfMediaInsightsPipelineConfiguration#keywords}.
	// Experimental.
	Keywords *[]*string `field:"required" json:"keywords" yaml:"keywords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#rule_name TfMediaInsightsPipelineConfiguration#rule_name}.
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#negate TfMediaInsightsPipelineConfiguration#negate}.
	// Experimental.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
}


package chimesdkmediapipelines


// Experimental.
type AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#keywords AwsMediaInsightsPipelineConfiguration#keywords}.
	// Experimental.
	Keywords *[]*string `field:"required" json:"keywords" yaml:"keywords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#rule_name AwsMediaInsightsPipelineConfiguration#rule_name}.
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#negate AwsMediaInsightsPipelineConfiguration#negate}.
	// Experimental.
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
}


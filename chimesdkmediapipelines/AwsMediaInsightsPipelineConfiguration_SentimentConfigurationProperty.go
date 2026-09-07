package chimesdkmediapipelines


// Experimental.
type AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#rule_name AwsMediaInsightsPipelineConfiguration#rule_name}.
	// Experimental.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sentiment_type AwsMediaInsightsPipelineConfiguration#sentiment_type}.
	// Experimental.
	SentimentType *string `field:"required" json:"sentimentType" yaml:"sentimentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#time_period AwsMediaInsightsPipelineConfiguration#time_period}.
	// Experimental.
	TimePeriod *float64 `field:"required" json:"timePeriod" yaml:"timePeriod"`
}


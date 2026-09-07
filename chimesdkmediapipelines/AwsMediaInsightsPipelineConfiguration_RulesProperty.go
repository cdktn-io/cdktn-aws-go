package chimesdkmediapipelines


// Experimental.
type AwsMediaInsightsPipelineConfiguration_RulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#type AwsMediaInsightsPipelineConfiguration#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// issue_detection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#issue_detection_configuration AwsMediaInsightsPipelineConfiguration#issue_detection_configuration}
	// Experimental.
	IssueDetectionConfiguration *AwsMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty `field:"optional" json:"issueDetectionConfiguration" yaml:"issueDetectionConfiguration"`
	// keyword_match_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#keyword_match_configuration AwsMediaInsightsPipelineConfiguration#keyword_match_configuration}
	// Experimental.
	KeywordMatchConfiguration *AwsMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty `field:"optional" json:"keywordMatchConfiguration" yaml:"keywordMatchConfiguration"`
	// sentiment_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sentiment_configuration AwsMediaInsightsPipelineConfiguration#sentiment_configuration}
	// Experimental.
	SentimentConfiguration *AwsMediaInsightsPipelineConfiguration_SentimentConfigurationProperty `field:"optional" json:"sentimentConfiguration" yaml:"sentimentConfiguration"`
}


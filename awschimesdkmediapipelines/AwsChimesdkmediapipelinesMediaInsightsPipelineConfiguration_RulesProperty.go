package awschimesdkmediapipelines


// Experimental.
type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_RulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// issue_detection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#issue_detection_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#issue_detection_configuration}
	// Experimental.
	IssueDetectionConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty `field:"optional" json:"issueDetectionConfiguration" yaml:"issueDetectionConfiguration"`
	// keyword_match_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#keyword_match_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#keyword_match_configuration}
	// Experimental.
	KeywordMatchConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty `field:"optional" json:"keywordMatchConfiguration" yaml:"keywordMatchConfiguration"`
	// sentiment_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sentiment_configuration AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#sentiment_configuration}
	// Experimental.
	SentimentConfiguration *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_SentimentConfigurationProperty `field:"optional" json:"sentimentConfiguration" yaml:"sentimentConfiguration"`
}


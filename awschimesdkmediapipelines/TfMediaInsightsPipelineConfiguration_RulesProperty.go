package awschimesdkmediapipelines


// Experimental.
type TfMediaInsightsPipelineConfiguration_RulesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#type TfMediaInsightsPipelineConfiguration#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// issue_detection_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#issue_detection_configuration TfMediaInsightsPipelineConfiguration#issue_detection_configuration}
	// Experimental.
	IssueDetectionConfiguration *TfMediaInsightsPipelineConfiguration_IssueDetectionConfigurationProperty `field:"optional" json:"issueDetectionConfiguration" yaml:"issueDetectionConfiguration"`
	// keyword_match_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#keyword_match_configuration TfMediaInsightsPipelineConfiguration#keyword_match_configuration}
	// Experimental.
	KeywordMatchConfiguration *TfMediaInsightsPipelineConfiguration_KeywordMatchConfigurationProperty `field:"optional" json:"keywordMatchConfiguration" yaml:"keywordMatchConfiguration"`
	// sentiment_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#sentiment_configuration TfMediaInsightsPipelineConfiguration#sentiment_configuration}
	// Experimental.
	SentimentConfiguration *TfMediaInsightsPipelineConfiguration_SentimentConfigurationProperty `field:"optional" json:"sentimentConfiguration" yaml:"sentimentConfiguration"`
}


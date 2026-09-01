package awschimesdkmediapipelines


// Experimental.
type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_code AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#language_code}.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#call_analytics_stream_categories AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#call_analytics_stream_categories}.
	// Experimental.
	CallAnalyticsStreamCategories *[]*string `field:"optional" json:"callAnalyticsStreamCategories" yaml:"callAnalyticsStreamCategories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_identification_type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#content_identification_type}.
	// Experimental.
	ContentIdentificationType *string `field:"optional" json:"contentIdentificationType" yaml:"contentIdentificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_redaction_type AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#content_redaction_type}.
	// Experimental.
	ContentRedactionType *string `field:"optional" json:"contentRedactionType" yaml:"contentRedactionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#enable_partial_results_stabilization AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#enable_partial_results_stabilization}.
	// Experimental.
	EnablePartialResultsStabilization interface{} `field:"optional" json:"enablePartialResultsStabilization" yaml:"enablePartialResultsStabilization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#filter_partial_results AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#filter_partial_results}.
	// Experimental.
	FilterPartialResults interface{} `field:"optional" json:"filterPartialResults" yaml:"filterPartialResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_model_name AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#language_model_name}.
	// Experimental.
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#partial_results_stability AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#partial_results_stability}.
	// Experimental.
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#pii_entity_types AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#pii_entity_types}.
	// Experimental.
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// post_call_analytics_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#post_call_analytics_settings AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#post_call_analytics_settings}
	// Experimental.
	PostCallAnalyticsSettings *AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty `field:"optional" json:"postCallAnalyticsSettings" yaml:"postCallAnalyticsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_method AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#vocabulary_filter_method}.
	// Experimental.
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_name AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#vocabulary_filter_name}.
	// Experimental.
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_name AwsChimesdkmediapipelinesMediaInsightsPipelineConfiguration#vocabulary_name}.
	// Experimental.
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}


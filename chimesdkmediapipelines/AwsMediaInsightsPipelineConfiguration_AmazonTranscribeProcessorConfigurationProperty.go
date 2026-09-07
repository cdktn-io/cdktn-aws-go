package chimesdkmediapipelines


// Experimental.
type AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_code AwsMediaInsightsPipelineConfiguration#language_code}.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_identification_type AwsMediaInsightsPipelineConfiguration#content_identification_type}.
	// Experimental.
	ContentIdentificationType *string `field:"optional" json:"contentIdentificationType" yaml:"contentIdentificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_redaction_type AwsMediaInsightsPipelineConfiguration#content_redaction_type}.
	// Experimental.
	ContentRedactionType *string `field:"optional" json:"contentRedactionType" yaml:"contentRedactionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#enable_partial_results_stabilization AwsMediaInsightsPipelineConfiguration#enable_partial_results_stabilization}.
	// Experimental.
	EnablePartialResultsStabilization interface{} `field:"optional" json:"enablePartialResultsStabilization" yaml:"enablePartialResultsStabilization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#filter_partial_results AwsMediaInsightsPipelineConfiguration#filter_partial_results}.
	// Experimental.
	FilterPartialResults interface{} `field:"optional" json:"filterPartialResults" yaml:"filterPartialResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_model_name AwsMediaInsightsPipelineConfiguration#language_model_name}.
	// Experimental.
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#partial_results_stability AwsMediaInsightsPipelineConfiguration#partial_results_stability}.
	// Experimental.
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#pii_entity_types AwsMediaInsightsPipelineConfiguration#pii_entity_types}.
	// Experimental.
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#show_speaker_label AwsMediaInsightsPipelineConfiguration#show_speaker_label}.
	// Experimental.
	ShowSpeakerLabel interface{} `field:"optional" json:"showSpeakerLabel" yaml:"showSpeakerLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_method AwsMediaInsightsPipelineConfiguration#vocabulary_filter_method}.
	// Experimental.
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_name AwsMediaInsightsPipelineConfiguration#vocabulary_filter_name}.
	// Experimental.
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_name AwsMediaInsightsPipelineConfiguration#vocabulary_name}.
	// Experimental.
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}


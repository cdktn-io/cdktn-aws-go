package awschimesdkmediapipelines


// Experimental.
type TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_code TfMediaInsightsPipelineConfiguration#language_code}.
	// Experimental.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_identification_type TfMediaInsightsPipelineConfiguration#content_identification_type}.
	// Experimental.
	ContentIdentificationType *string `field:"optional" json:"contentIdentificationType" yaml:"contentIdentificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_redaction_type TfMediaInsightsPipelineConfiguration#content_redaction_type}.
	// Experimental.
	ContentRedactionType *string `field:"optional" json:"contentRedactionType" yaml:"contentRedactionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#enable_partial_results_stabilization TfMediaInsightsPipelineConfiguration#enable_partial_results_stabilization}.
	// Experimental.
	EnablePartialResultsStabilization interface{} `field:"optional" json:"enablePartialResultsStabilization" yaml:"enablePartialResultsStabilization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#filter_partial_results TfMediaInsightsPipelineConfiguration#filter_partial_results}.
	// Experimental.
	FilterPartialResults interface{} `field:"optional" json:"filterPartialResults" yaml:"filterPartialResults"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#language_model_name TfMediaInsightsPipelineConfiguration#language_model_name}.
	// Experimental.
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#partial_results_stability TfMediaInsightsPipelineConfiguration#partial_results_stability}.
	// Experimental.
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#pii_entity_types TfMediaInsightsPipelineConfiguration#pii_entity_types}.
	// Experimental.
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#show_speaker_label TfMediaInsightsPipelineConfiguration#show_speaker_label}.
	// Experimental.
	ShowSpeakerLabel interface{} `field:"optional" json:"showSpeakerLabel" yaml:"showSpeakerLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_method TfMediaInsightsPipelineConfiguration#vocabulary_filter_method}.
	// Experimental.
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_filter_name TfMediaInsightsPipelineConfiguration#vocabulary_filter_name}.
	// Experimental.
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#vocabulary_name TfMediaInsightsPipelineConfiguration#vocabulary_name}.
	// Experimental.
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}


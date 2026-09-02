package awschimesdkmediapipelines


// Experimental.
type TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#speaker_search_status TfMediaInsightsPipelineConfiguration#speaker_search_status}.
	// Experimental.
	SpeakerSearchStatus *string `field:"required" json:"speakerSearchStatus" yaml:"speakerSearchStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#voice_tone_analysis_status TfMediaInsightsPipelineConfiguration#voice_tone_analysis_status}.
	// Experimental.
	VoiceToneAnalysisStatus *string `field:"required" json:"voiceToneAnalysisStatus" yaml:"voiceToneAnalysisStatus"`
}


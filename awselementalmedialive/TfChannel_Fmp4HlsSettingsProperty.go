package awselementalmedialive


// Experimental.
type TfChannel_Fmp4HlsSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_rendition_sets TfChannel#audio_rendition_sets}.
	// Experimental.
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_id3_behavior TfChannel#nielsen_id3_behavior}.
	// Experimental.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timed_metadata_behavior TfChannel#timed_metadata_behavior}.
	// Experimental.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
}


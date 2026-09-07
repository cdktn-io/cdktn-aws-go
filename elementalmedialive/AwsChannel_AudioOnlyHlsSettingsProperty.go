package elementalmedialive


// Experimental.
type AwsChannel_AudioOnlyHlsSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_group_id AwsChannel#audio_group_id}.
	// Experimental.
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// audio_only_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_only_image AwsChannel#audio_only_image}
	// Experimental.
	AudioOnlyImage *AwsChannel_AudioOnlyImageProperty `field:"optional" json:"audioOnlyImage" yaml:"audioOnlyImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_track_type AwsChannel#audio_track_type}.
	// Experimental.
	AudioTrackType *string `field:"optional" json:"audioTrackType" yaml:"audioTrackType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segment_type AwsChannel#segment_type}.
	// Experimental.
	SegmentType *string `field:"optional" json:"segmentType" yaml:"segmentType"`
}


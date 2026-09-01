package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_AudioOnlyHlsSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_group_id AwsMedialiveChannel#audio_group_id}.
	// Experimental.
	AudioGroupId *string `field:"optional" json:"audioGroupId" yaml:"audioGroupId"`
	// audio_only_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_only_image AwsMedialiveChannel#audio_only_image}
	// Experimental.
	AudioOnlyImage *AwsMedialiveChannel_AudioOnlyImageProperty `field:"optional" json:"audioOnlyImage" yaml:"audioOnlyImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_track_type AwsMedialiveChannel#audio_track_type}.
	// Experimental.
	AudioTrackType *string `field:"optional" json:"audioTrackType" yaml:"audioTrackType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segment_type AwsMedialiveChannel#segment_type}.
	// Experimental.
	SegmentType *string `field:"optional" json:"segmentType" yaml:"segmentType"`
}


package elementalmedialive


// Experimental.
type AwsChannel_MsSmoothGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsChannel#destination}
	// Experimental.
	Destination *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#acquisition_point_id AwsChannel#acquisition_point_id}.
	// Experimental.
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_only_timecode_control AwsChannel#audio_only_timecode_control}.
	// Experimental.
	AudioOnlyTimecodeControl *string `field:"optional" json:"audioOnlyTimecodeControl" yaml:"audioOnlyTimecodeControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#certificate_mode AwsChannel#certificate_mode}.
	// Experimental.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#connection_retry_interval AwsChannel#connection_retry_interval}.
	// Experimental.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#event_id AwsChannel#event_id}.
	// Experimental.
	EventId *string `field:"optional" json:"eventId" yaml:"eventId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#event_id_mode AwsChannel#event_id_mode}.
	// Experimental.
	EventIdMode *string `field:"optional" json:"eventIdMode" yaml:"eventIdMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#event_stop_behavior AwsChannel#event_stop_behavior}.
	// Experimental.
	EventStopBehavior *string `field:"optional" json:"eventStopBehavior" yaml:"eventStopBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#filecache_duration AwsChannel#filecache_duration}.
	// Experimental.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fragment_length AwsChannel#fragment_length}.
	// Experimental.
	FragmentLength *float64 `field:"optional" json:"fragmentLength" yaml:"fragmentLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_action AwsChannel#input_loss_action}.
	// Experimental.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#num_retries AwsChannel#num_retries}.
	// Experimental.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#restart_delay AwsChannel#restart_delay}.
	// Experimental.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segmentation_mode AwsChannel#segmentation_mode}.
	// Experimental.
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#send_delay_ms AwsChannel#send_delay_ms}.
	// Experimental.
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sparse_track_type AwsChannel#sparse_track_type}.
	// Experimental.
	SparseTrackType *string `field:"optional" json:"sparseTrackType" yaml:"sparseTrackType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#stream_manifest_behavior AwsChannel#stream_manifest_behavior}.
	// Experimental.
	StreamManifestBehavior *string `field:"optional" json:"streamManifestBehavior" yaml:"streamManifestBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timestamp_offset AwsChannel#timestamp_offset}.
	// Experimental.
	TimestampOffset *string `field:"optional" json:"timestampOffset" yaml:"timestampOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timestamp_offset_mode AwsChannel#timestamp_offset_mode}.
	// Experimental.
	TimestampOffsetMode *string `field:"optional" json:"timestampOffsetMode" yaml:"timestampOffsetMode"`
}


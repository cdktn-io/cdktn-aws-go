package awselementalmedialive


// Experimental.
type TfChannel_MsSmoothGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination TfChannel#destination}
	// Experimental.
	Destination *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#acquisition_point_id TfChannel#acquisition_point_id}.
	// Experimental.
	AcquisitionPointId *string `field:"optional" json:"acquisitionPointId" yaml:"acquisitionPointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_only_timecode_control TfChannel#audio_only_timecode_control}.
	// Experimental.
	AudioOnlyTimecodeControl *string `field:"optional" json:"audioOnlyTimecodeControl" yaml:"audioOnlyTimecodeControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#certificate_mode TfChannel#certificate_mode}.
	// Experimental.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#connection_retry_interval TfChannel#connection_retry_interval}.
	// Experimental.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#event_id TfChannel#event_id}.
	// Experimental.
	EventId *string `field:"optional" json:"eventId" yaml:"eventId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#event_id_mode TfChannel#event_id_mode}.
	// Experimental.
	EventIdMode *string `field:"optional" json:"eventIdMode" yaml:"eventIdMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#event_stop_behavior TfChannel#event_stop_behavior}.
	// Experimental.
	EventStopBehavior *string `field:"optional" json:"eventStopBehavior" yaml:"eventStopBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#filecache_duration TfChannel#filecache_duration}.
	// Experimental.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fragment_length TfChannel#fragment_length}.
	// Experimental.
	FragmentLength *float64 `field:"optional" json:"fragmentLength" yaml:"fragmentLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_action TfChannel#input_loss_action}.
	// Experimental.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#num_retries TfChannel#num_retries}.
	// Experimental.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#restart_delay TfChannel#restart_delay}.
	// Experimental.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segmentation_mode TfChannel#segmentation_mode}.
	// Experimental.
	SegmentationMode *string `field:"optional" json:"segmentationMode" yaml:"segmentationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#send_delay_ms TfChannel#send_delay_ms}.
	// Experimental.
	SendDelayMs *float64 `field:"optional" json:"sendDelayMs" yaml:"sendDelayMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#sparse_track_type TfChannel#sparse_track_type}.
	// Experimental.
	SparseTrackType *string `field:"optional" json:"sparseTrackType" yaml:"sparseTrackType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#stream_manifest_behavior TfChannel#stream_manifest_behavior}.
	// Experimental.
	StreamManifestBehavior *string `field:"optional" json:"streamManifestBehavior" yaml:"streamManifestBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timestamp_offset TfChannel#timestamp_offset}.
	// Experimental.
	TimestampOffset *string `field:"optional" json:"timestampOffset" yaml:"timestampOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timestamp_offset_mode TfChannel#timestamp_offset_mode}.
	// Experimental.
	TimestampOffsetMode *string `field:"optional" json:"timestampOffsetMode" yaml:"timestampOffsetMode"`
}


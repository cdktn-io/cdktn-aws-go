package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#absent_input_audio_behavior TfChannel#absent_input_audio_behavior}.
	// Experimental.
	AbsentInputAudioBehavior *string `field:"optional" json:"absentInputAudioBehavior" yaml:"absentInputAudioBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#arib TfChannel#arib}.
	// Experimental.
	Arib *string `field:"optional" json:"arib" yaml:"arib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#arib_captions_pid TfChannel#arib_captions_pid}.
	// Experimental.
	AribCaptionsPid *string `field:"optional" json:"aribCaptionsPid" yaml:"aribCaptionsPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#arib_captions_pid_control TfChannel#arib_captions_pid_control}.
	// Experimental.
	AribCaptionsPidControl *string `field:"optional" json:"aribCaptionsPidControl" yaml:"aribCaptionsPidControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_buffer_model TfChannel#audio_buffer_model}.
	// Experimental.
	AudioBufferModel *string `field:"optional" json:"audioBufferModel" yaml:"audioBufferModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_frames_per_pes TfChannel#audio_frames_per_pes}.
	// Experimental.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_pids TfChannel#audio_pids}.
	// Experimental.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_stream_type TfChannel#audio_stream_type}.
	// Experimental.
	AudioStreamType *string `field:"optional" json:"audioStreamType" yaml:"audioStreamType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitrate TfChannel#bitrate}.
	// Experimental.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buffer_model TfChannel#buffer_model}.
	// Experimental.
	BufferModel *string `field:"optional" json:"bufferModel" yaml:"bufferModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#cc_descriptor TfChannel#cc_descriptor}.
	// Experimental.
	CcDescriptor *string `field:"optional" json:"ccDescriptor" yaml:"ccDescriptor"`
	// dvb_nit_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_nit_settings TfChannel#dvb_nit_settings}
	// Experimental.
	DvbNitSettings *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbNitSettingsProperty `field:"optional" json:"dvbNitSettings" yaml:"dvbNitSettings"`
	// dvb_sdt_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_sdt_settings TfChannel#dvb_sdt_settings}
	// Experimental.
	DvbSdtSettings *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbSdtSettingsProperty `field:"optional" json:"dvbSdtSettings" yaml:"dvbSdtSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_sub_pids TfChannel#dvb_sub_pids}.
	// Experimental.
	DvbSubPids *string `field:"optional" json:"dvbSubPids" yaml:"dvbSubPids"`
	// dvb_tdt_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_tdt_settings TfChannel#dvb_tdt_settings}
	// Experimental.
	DvbTdtSettings *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsUdpOutputSettingsContainerSettingsM2tsSettingsDvbTdtSettingsProperty `field:"optional" json:"dvbTdtSettings" yaml:"dvbTdtSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#dvb_teletext_pid TfChannel#dvb_teletext_pid}.
	// Experimental.
	DvbTeletextPid *string `field:"optional" json:"dvbTeletextPid" yaml:"dvbTeletextPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ebif TfChannel#ebif}.
	// Experimental.
	Ebif *string `field:"optional" json:"ebif" yaml:"ebif"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ebp_audio_interval TfChannel#ebp_audio_interval}.
	// Experimental.
	EbpAudioInterval *string `field:"optional" json:"ebpAudioInterval" yaml:"ebpAudioInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ebp_lookahead_ms TfChannel#ebp_lookahead_ms}.
	// Experimental.
	EbpLookaheadMs *float64 `field:"optional" json:"ebpLookaheadMs" yaml:"ebpLookaheadMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ebp_placement TfChannel#ebp_placement}.
	// Experimental.
	EbpPlacement *string `field:"optional" json:"ebpPlacement" yaml:"ebpPlacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ecm_pid TfChannel#ecm_pid}.
	// Experimental.
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#es_rate_in_pes TfChannel#es_rate_in_pes}.
	// Experimental.
	EsRateInPes *string `field:"optional" json:"esRateInPes" yaml:"esRateInPes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#etv_platform_pid TfChannel#etv_platform_pid}.
	// Experimental.
	EtvPlatformPid *string `field:"optional" json:"etvPlatformPid" yaml:"etvPlatformPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#etv_signal_pid TfChannel#etv_signal_pid}.
	// Experimental.
	EtvSignalPid *string `field:"optional" json:"etvSignalPid" yaml:"etvSignalPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fragment_time TfChannel#fragment_time}.
	// Experimental.
	FragmentTime *float64 `field:"optional" json:"fragmentTime" yaml:"fragmentTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#klv TfChannel#klv}.
	// Experimental.
	Klv *string `field:"optional" json:"klv" yaml:"klv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#klv_data_pids TfChannel#klv_data_pids}.
	// Experimental.
	KlvDataPids *string `field:"optional" json:"klvDataPids" yaml:"klvDataPids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_id3_behavior TfChannel#nielsen_id3_behavior}.
	// Experimental.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#null_packet_bitrate TfChannel#null_packet_bitrate}.
	// Experimental.
	NullPacketBitrate *float64 `field:"optional" json:"nullPacketBitrate" yaml:"nullPacketBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pat_interval TfChannel#pat_interval}.
	// Experimental.
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pcr_control TfChannel#pcr_control}.
	// Experimental.
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pcr_period TfChannel#pcr_period}.
	// Experimental.
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pcr_pid TfChannel#pcr_pid}.
	// Experimental.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pmt_interval TfChannel#pmt_interval}.
	// Experimental.
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#pmt_pid TfChannel#pmt_pid}.
	// Experimental.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#program_num TfChannel#program_num}.
	// Experimental.
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rate_mode TfChannel#rate_mode}.
	// Experimental.
	RateMode *string `field:"optional" json:"rateMode" yaml:"rateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte27_pids TfChannel#scte27_pids}.
	// Experimental.
	Scte27Pids *string `field:"optional" json:"scte27Pids" yaml:"scte27Pids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte35_control TfChannel#scte35_control}.
	// Experimental.
	Scte35Control *string `field:"optional" json:"scte35Control" yaml:"scte35Control"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scte35_pid TfChannel#scte35_pid}.
	// Experimental.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segmentation_markers TfChannel#segmentation_markers}.
	// Experimental.
	SegmentationMarkers *string `field:"optional" json:"segmentationMarkers" yaml:"segmentationMarkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segmentation_style TfChannel#segmentation_style}.
	// Experimental.
	SegmentationStyle *string `field:"optional" json:"segmentationStyle" yaml:"segmentationStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segmentation_time TfChannel#segmentation_time}.
	// Experimental.
	SegmentationTime *float64 `field:"optional" json:"segmentationTime" yaml:"segmentationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timed_metadata_behavior TfChannel#timed_metadata_behavior}.
	// Experimental.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timed_metadata_pid TfChannel#timed_metadata_pid}.
	// Experimental.
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#transport_stream_id TfChannel#transport_stream_id}.
	// Experimental.
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_pid TfChannel#video_pid}.
	// Experimental.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}


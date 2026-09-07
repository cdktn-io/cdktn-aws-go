package elementalmedialive


// Experimental.
type AwsChannel_HlsGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsChannel#destination}
	// Experimental.
	Destination *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ad_markers AwsChannel#ad_markers}.
	// Experimental.
	AdMarkers *[]*string `field:"optional" json:"adMarkers" yaml:"adMarkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#base_url_content AwsChannel#base_url_content}.
	// Experimental.
	BaseUrlContent *string `field:"optional" json:"baseUrlContent" yaml:"baseUrlContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#base_url_content1 AwsChannel#base_url_content1}.
	// Experimental.
	BaseUrlContent1 *string `field:"optional" json:"baseUrlContent1" yaml:"baseUrlContent1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#base_url_manifest AwsChannel#base_url_manifest}.
	// Experimental.
	BaseUrlManifest *string `field:"optional" json:"baseUrlManifest" yaml:"baseUrlManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#base_url_manifest1 AwsChannel#base_url_manifest1}.
	// Experimental.
	BaseUrlManifest1 *string `field:"optional" json:"baseUrlManifest1" yaml:"baseUrlManifest1"`
	// caption_language_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_language_mappings AwsChannel#caption_language_mappings}
	// Experimental.
	CaptionLanguageMappings interface{} `field:"optional" json:"captionLanguageMappings" yaml:"captionLanguageMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_language_setting AwsChannel#caption_language_setting}.
	// Experimental.
	CaptionLanguageSetting *string `field:"optional" json:"captionLanguageSetting" yaml:"captionLanguageSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#client_cache AwsChannel#client_cache}.
	// Experimental.
	ClientCache *string `field:"optional" json:"clientCache" yaml:"clientCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#codec_specification AwsChannel#codec_specification}.
	// Experimental.
	CodecSpecification *string `field:"optional" json:"codecSpecification" yaml:"codecSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#constant_iv AwsChannel#constant_iv}.
	// Experimental.
	ConstantIv *string `field:"optional" json:"constantIv" yaml:"constantIv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#directory_structure AwsChannel#directory_structure}.
	// Experimental.
	DirectoryStructure *string `field:"optional" json:"directoryStructure" yaml:"directoryStructure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#discontinuity_tags AwsChannel#discontinuity_tags}.
	// Experimental.
	DiscontinuityTags *string `field:"optional" json:"discontinuityTags" yaml:"discontinuityTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#encryption_type AwsChannel#encryption_type}.
	// Experimental.
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// hls_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_cdn_settings AwsChannel#hls_cdn_settings}
	// Experimental.
	HlsCdnSettings interface{} `field:"optional" json:"hlsCdnSettings" yaml:"hlsCdnSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_id3_segment_tagging AwsChannel#hls_id3_segment_tagging}.
	// Experimental.
	HlsId3SegmentTagging *string `field:"optional" json:"hlsId3SegmentTagging" yaml:"hlsId3SegmentTagging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#iframe_only_playlists AwsChannel#iframe_only_playlists}.
	// Experimental.
	IframeOnlyPlaylists *string `field:"optional" json:"iframeOnlyPlaylists" yaml:"iframeOnlyPlaylists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#incomplete_segment_behavior AwsChannel#incomplete_segment_behavior}.
	// Experimental.
	IncompleteSegmentBehavior *string `field:"optional" json:"incompleteSegmentBehavior" yaml:"incompleteSegmentBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#index_n_segments AwsChannel#index_n_segments}.
	// Experimental.
	IndexNSegments *float64 `field:"optional" json:"indexNSegments" yaml:"indexNSegments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#input_loss_action AwsChannel#input_loss_action}.
	// Experimental.
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#iv_in_manifest AwsChannel#iv_in_manifest}.
	// Experimental.
	IvInManifest *string `field:"optional" json:"ivInManifest" yaml:"ivInManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#iv_source AwsChannel#iv_source}.
	// Experimental.
	IvSource *string `field:"optional" json:"ivSource" yaml:"ivSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#keep_segments AwsChannel#keep_segments}.
	// Experimental.
	KeepSegments *float64 `field:"optional" json:"keepSegments" yaml:"keepSegments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#key_format AwsChannel#key_format}.
	// Experimental.
	KeyFormat *string `field:"optional" json:"keyFormat" yaml:"keyFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#key_format_versions AwsChannel#key_format_versions}.
	// Experimental.
	KeyFormatVersions *string `field:"optional" json:"keyFormatVersions" yaml:"keyFormatVersions"`
	// key_provider_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#key_provider_settings AwsChannel#key_provider_settings}
	// Experimental.
	KeyProviderSettings *AwsChannel_KeyProviderSettingsProperty `field:"optional" json:"keyProviderSettings" yaml:"keyProviderSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#manifest_compression AwsChannel#manifest_compression}.
	// Experimental.
	ManifestCompression *string `field:"optional" json:"manifestCompression" yaml:"manifestCompression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#manifest_duration_format AwsChannel#manifest_duration_format}.
	// Experimental.
	ManifestDurationFormat *string `field:"optional" json:"manifestDurationFormat" yaml:"manifestDurationFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#min_segment_length AwsChannel#min_segment_length}.
	// Experimental.
	MinSegmentLength *float64 `field:"optional" json:"minSegmentLength" yaml:"minSegmentLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#mode AwsChannel#mode}.
	// Experimental.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_selection AwsChannel#output_selection}.
	// Experimental.
	OutputSelection *string `field:"optional" json:"outputSelection" yaml:"outputSelection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#program_date_time AwsChannel#program_date_time}.
	// Experimental.
	ProgramDateTime *string `field:"optional" json:"programDateTime" yaml:"programDateTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#program_date_time_clock AwsChannel#program_date_time_clock}.
	// Experimental.
	ProgramDateTimeClock *string `field:"optional" json:"programDateTimeClock" yaml:"programDateTimeClock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#program_date_time_period AwsChannel#program_date_time_period}.
	// Experimental.
	ProgramDateTimePeriod *float64 `field:"optional" json:"programDateTimePeriod" yaml:"programDateTimePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#redundant_manifest AwsChannel#redundant_manifest}.
	// Experimental.
	RedundantManifest *string `field:"optional" json:"redundantManifest" yaml:"redundantManifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segment_length AwsChannel#segment_length}.
	// Experimental.
	SegmentLength *float64 `field:"optional" json:"segmentLength" yaml:"segmentLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#segments_per_subdirectory AwsChannel#segments_per_subdirectory}.
	// Experimental.
	SegmentsPerSubdirectory *float64 `field:"optional" json:"segmentsPerSubdirectory" yaml:"segmentsPerSubdirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#stream_inf_resolution AwsChannel#stream_inf_resolution}.
	// Experimental.
	StreamInfResolution *string `field:"optional" json:"streamInfResolution" yaml:"streamInfResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timed_metadata_id3_frame AwsChannel#timed_metadata_id3_frame}.
	// Experimental.
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timed_metadata_id3_period AwsChannel#timed_metadata_id3_period}.
	// Experimental.
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timestamp_delta_milliseconds AwsChannel#timestamp_delta_milliseconds}.
	// Experimental.
	TimestampDeltaMilliseconds *float64 `field:"optional" json:"timestampDeltaMilliseconds" yaml:"timestampDeltaMilliseconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#ts_file_mode AwsChannel#ts_file_mode}.
	// Experimental.
	TsFileMode *string `field:"optional" json:"tsFileMode" yaml:"tsFileMode"`
}


package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_H264SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#adaptive_quantization AwsMedialiveChannel#adaptive_quantization}.
	// Experimental.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#afd_signaling AwsMedialiveChannel#afd_signaling}.
	// Experimental.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitrate AwsMedialiveChannel#bitrate}.
	// Experimental.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buf_fill_pct AwsMedialiveChannel#buf_fill_pct}.
	// Experimental.
	BufFillPct *float64 `field:"optional" json:"bufFillPct" yaml:"bufFillPct"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buf_size AwsMedialiveChannel#buf_size}.
	// Experimental.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#color_metadata AwsMedialiveChannel#color_metadata}.
	// Experimental.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#entropy_encoding AwsMedialiveChannel#entropy_encoding}.
	// Experimental.
	EntropyEncoding *string `field:"optional" json:"entropyEncoding" yaml:"entropyEncoding"`
	// filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#filter_settings AwsMedialiveChannel#filter_settings}
	// Experimental.
	FilterSettings *AwsMedialiveChannel_EncoderSettingsVideoDescriptionsCodecSettingsH264SettingsFilterSettingsProperty `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fixed_afd AwsMedialiveChannel#fixed_afd}.
	// Experimental.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#flicker_aq AwsMedialiveChannel#flicker_aq}.
	// Experimental.
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#force_field_pictures AwsMedialiveChannel#force_field_pictures}.
	// Experimental.
	ForceFieldPictures *string `field:"optional" json:"forceFieldPictures" yaml:"forceFieldPictures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#framerate_control AwsMedialiveChannel#framerate_control}.
	// Experimental.
	FramerateControl *string `field:"optional" json:"framerateControl" yaml:"framerateControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#framerate_denominator AwsMedialiveChannel#framerate_denominator}.
	// Experimental.
	FramerateDenominator *float64 `field:"optional" json:"framerateDenominator" yaml:"framerateDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#framerate_numerator AwsMedialiveChannel#framerate_numerator}.
	// Experimental.
	FramerateNumerator *float64 `field:"optional" json:"framerateNumerator" yaml:"framerateNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_b_reference AwsMedialiveChannel#gop_b_reference}.
	// Experimental.
	GopBReference *string `field:"optional" json:"gopBReference" yaml:"gopBReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_closed_cadence AwsMedialiveChannel#gop_closed_cadence}.
	// Experimental.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_num_b_frames AwsMedialiveChannel#gop_num_b_frames}.
	// Experimental.
	GopNumBFrames *float64 `field:"optional" json:"gopNumBFrames" yaml:"gopNumBFrames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_size AwsMedialiveChannel#gop_size}.
	// Experimental.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_size_units AwsMedialiveChannel#gop_size_units}.
	// Experimental.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#level AwsMedialiveChannel#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#look_ahead_rate_control AwsMedialiveChannel#look_ahead_rate_control}.
	// Experimental.
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#max_bitrate AwsMedialiveChannel#max_bitrate}.
	// Experimental.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#min_i_interval AwsMedialiveChannel#min_i_interval}.
	// Experimental.
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#num_ref_frames AwsMedialiveChannel#num_ref_frames}.
	// Experimental.
	NumRefFrames *float64 `field:"optional" json:"numRefFrames" yaml:"numRefFrames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#par_control AwsMedialiveChannel#par_control}.
	// Experimental.
	ParControl *string `field:"optional" json:"parControl" yaml:"parControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#par_denominator AwsMedialiveChannel#par_denominator}.
	// Experimental.
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#par_numerator AwsMedialiveChannel#par_numerator}.
	// Experimental.
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#profile AwsMedialiveChannel#profile}.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#quality_level AwsMedialiveChannel#quality_level}.
	// Experimental.
	QualityLevel *string `field:"optional" json:"qualityLevel" yaml:"qualityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#qvbr_quality_level AwsMedialiveChannel#qvbr_quality_level}.
	// Experimental.
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rate_control_mode AwsMedialiveChannel#rate_control_mode}.
	// Experimental.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scan_type AwsMedialiveChannel#scan_type}.
	// Experimental.
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scene_change_detect AwsMedialiveChannel#scene_change_detect}.
	// Experimental.
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#slices AwsMedialiveChannel#slices}.
	// Experimental.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#softness AwsMedialiveChannel#softness}.
	// Experimental.
	Softness *float64 `field:"optional" json:"softness" yaml:"softness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#spatial_aq AwsMedialiveChannel#spatial_aq}.
	// Experimental.
	SpatialAq *string `field:"optional" json:"spatialAq" yaml:"spatialAq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#subgop_length AwsMedialiveChannel#subgop_length}.
	// Experimental.
	SubgopLength *string `field:"optional" json:"subgopLength" yaml:"subgopLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#syntax AwsMedialiveChannel#syntax}.
	// Experimental.
	Syntax *string `field:"optional" json:"syntax" yaml:"syntax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#temporal_aq AwsMedialiveChannel#temporal_aq}.
	// Experimental.
	TemporalAq *string `field:"optional" json:"temporalAq" yaml:"temporalAq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timecode_insertion AwsMedialiveChannel#timecode_insertion}.
	// Experimental.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
}


package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_H265SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitrate AwsMedialiveChannel#bitrate}.
	// Experimental.
	Bitrate *float64 `field:"required" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#framerate_denominator AwsMedialiveChannel#framerate_denominator}.
	// Experimental.
	FramerateDenominator *float64 `field:"required" json:"framerateDenominator" yaml:"framerateDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#framerate_numerator AwsMedialiveChannel#framerate_numerator}.
	// Experimental.
	FramerateNumerator *float64 `field:"required" json:"framerateNumerator" yaml:"framerateNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#adaptive_quantization AwsMedialiveChannel#adaptive_quantization}.
	// Experimental.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#afd_signaling AwsMedialiveChannel#afd_signaling}.
	// Experimental.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#alternative_transfer_function AwsMedialiveChannel#alternative_transfer_function}.
	// Experimental.
	AlternativeTransferFunction *string `field:"optional" json:"alternativeTransferFunction" yaml:"alternativeTransferFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buf_size AwsMedialiveChannel#buf_size}.
	// Experimental.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#color_metadata AwsMedialiveChannel#color_metadata}.
	// Experimental.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// color_space_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#color_space_settings AwsMedialiveChannel#color_space_settings}
	// Experimental.
	ColorSpaceSettings *AwsMedialiveChannel_ColorSpaceSettingsProperty `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#filter_settings AwsMedialiveChannel#filter_settings}
	// Experimental.
	FilterSettings *AwsMedialiveChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fixed_afd AwsMedialiveChannel#fixed_afd}.
	// Experimental.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#flicker_aq AwsMedialiveChannel#flicker_aq}.
	// Experimental.
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_closed_cadence AwsMedialiveChannel#gop_closed_cadence}.
	// Experimental.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#min_qp AwsMedialiveChannel#min_qp}.
	// Experimental.
	MinQp *float64 `field:"optional" json:"minQp" yaml:"minQp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#mv_over_picture_boundaries AwsMedialiveChannel#mv_over_picture_boundaries}.
	// Experimental.
	MvOverPictureBoundaries *string `field:"optional" json:"mvOverPictureBoundaries" yaml:"mvOverPictureBoundaries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#mv_temporal_predictor AwsMedialiveChannel#mv_temporal_predictor}.
	// Experimental.
	MvTemporalPredictor *string `field:"optional" json:"mvTemporalPredictor" yaml:"mvTemporalPredictor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#par_denominator AwsMedialiveChannel#par_denominator}.
	// Experimental.
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#par_numerator AwsMedialiveChannel#par_numerator}.
	// Experimental.
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#profile AwsMedialiveChannel#profile}.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tier AwsMedialiveChannel#tier}.
	// Experimental.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tile_height AwsMedialiveChannel#tile_height}.
	// Experimental.
	TileHeight *float64 `field:"optional" json:"tileHeight" yaml:"tileHeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tile_padding AwsMedialiveChannel#tile_padding}.
	// Experimental.
	TilePadding *string `field:"optional" json:"tilePadding" yaml:"tilePadding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tile_width AwsMedialiveChannel#tile_width}.
	// Experimental.
	TileWidth *float64 `field:"optional" json:"tileWidth" yaml:"tileWidth"`
	// timecode_burnin_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timecode_burnin_settings AwsMedialiveChannel#timecode_burnin_settings}
	// Experimental.
	TimecodeBurninSettings *AwsMedialiveChannel_TimecodeBurninSettingsProperty `field:"optional" json:"timecodeBurninSettings" yaml:"timecodeBurninSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timecode_insertion AwsMedialiveChannel#timecode_insertion}.
	// Experimental.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#treeblock_size AwsMedialiveChannel#treeblock_size}.
	// Experimental.
	TreeblockSize *string `field:"optional" json:"treeblockSize" yaml:"treeblockSize"`
}


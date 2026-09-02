package awselementalmedialive


// Experimental.
type TfChannel_H265SettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#bitrate TfChannel#bitrate}.
	// Experimental.
	Bitrate *float64 `field:"required" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#framerate_denominator TfChannel#framerate_denominator}.
	// Experimental.
	FramerateDenominator *float64 `field:"required" json:"framerateDenominator" yaml:"framerateDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#framerate_numerator TfChannel#framerate_numerator}.
	// Experimental.
	FramerateNumerator *float64 `field:"required" json:"framerateNumerator" yaml:"framerateNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#adaptive_quantization TfChannel#adaptive_quantization}.
	// Experimental.
	AdaptiveQuantization *string `field:"optional" json:"adaptiveQuantization" yaml:"adaptiveQuantization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#afd_signaling TfChannel#afd_signaling}.
	// Experimental.
	AfdSignaling *string `field:"optional" json:"afdSignaling" yaml:"afdSignaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#alternative_transfer_function TfChannel#alternative_transfer_function}.
	// Experimental.
	AlternativeTransferFunction *string `field:"optional" json:"alternativeTransferFunction" yaml:"alternativeTransferFunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#buf_size TfChannel#buf_size}.
	// Experimental.
	BufSize *float64 `field:"optional" json:"bufSize" yaml:"bufSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#color_metadata TfChannel#color_metadata}.
	// Experimental.
	ColorMetadata *string `field:"optional" json:"colorMetadata" yaml:"colorMetadata"`
	// color_space_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#color_space_settings TfChannel#color_space_settings}
	// Experimental.
	ColorSpaceSettings *TfChannel_ColorSpaceSettingsProperty `field:"optional" json:"colorSpaceSettings" yaml:"colorSpaceSettings"`
	// filter_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#filter_settings TfChannel#filter_settings}
	// Experimental.
	FilterSettings *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsH265SettingsFilterSettingsProperty `field:"optional" json:"filterSettings" yaml:"filterSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#fixed_afd TfChannel#fixed_afd}.
	// Experimental.
	FixedAfd *string `field:"optional" json:"fixedAfd" yaml:"fixedAfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#flicker_aq TfChannel#flicker_aq}.
	// Experimental.
	FlickerAq *string `field:"optional" json:"flickerAq" yaml:"flickerAq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_closed_cadence TfChannel#gop_closed_cadence}.
	// Experimental.
	GopClosedCadence *float64 `field:"optional" json:"gopClosedCadence" yaml:"gopClosedCadence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_size TfChannel#gop_size}.
	// Experimental.
	GopSize *float64 `field:"optional" json:"gopSize" yaml:"gopSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#gop_size_units TfChannel#gop_size_units}.
	// Experimental.
	GopSizeUnits *string `field:"optional" json:"gopSizeUnits" yaml:"gopSizeUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#level TfChannel#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#look_ahead_rate_control TfChannel#look_ahead_rate_control}.
	// Experimental.
	LookAheadRateControl *string `field:"optional" json:"lookAheadRateControl" yaml:"lookAheadRateControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#max_bitrate TfChannel#max_bitrate}.
	// Experimental.
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#min_i_interval TfChannel#min_i_interval}.
	// Experimental.
	MinIInterval *float64 `field:"optional" json:"minIInterval" yaml:"minIInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#min_qp TfChannel#min_qp}.
	// Experimental.
	MinQp *float64 `field:"optional" json:"minQp" yaml:"minQp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#mv_over_picture_boundaries TfChannel#mv_over_picture_boundaries}.
	// Experimental.
	MvOverPictureBoundaries *string `field:"optional" json:"mvOverPictureBoundaries" yaml:"mvOverPictureBoundaries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#mv_temporal_predictor TfChannel#mv_temporal_predictor}.
	// Experimental.
	MvTemporalPredictor *string `field:"optional" json:"mvTemporalPredictor" yaml:"mvTemporalPredictor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#par_denominator TfChannel#par_denominator}.
	// Experimental.
	ParDenominator *float64 `field:"optional" json:"parDenominator" yaml:"parDenominator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#par_numerator TfChannel#par_numerator}.
	// Experimental.
	ParNumerator *float64 `field:"optional" json:"parNumerator" yaml:"parNumerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#profile TfChannel#profile}.
	// Experimental.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#qvbr_quality_level TfChannel#qvbr_quality_level}.
	// Experimental.
	QvbrQualityLevel *float64 `field:"optional" json:"qvbrQualityLevel" yaml:"qvbrQualityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#rate_control_mode TfChannel#rate_control_mode}.
	// Experimental.
	RateControlMode *string `field:"optional" json:"rateControlMode" yaml:"rateControlMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scan_type TfChannel#scan_type}.
	// Experimental.
	ScanType *string `field:"optional" json:"scanType" yaml:"scanType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#scene_change_detect TfChannel#scene_change_detect}.
	// Experimental.
	SceneChangeDetect *string `field:"optional" json:"sceneChangeDetect" yaml:"sceneChangeDetect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#slices TfChannel#slices}.
	// Experimental.
	Slices *float64 `field:"optional" json:"slices" yaml:"slices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tier TfChannel#tier}.
	// Experimental.
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tile_height TfChannel#tile_height}.
	// Experimental.
	TileHeight *float64 `field:"optional" json:"tileHeight" yaml:"tileHeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tile_padding TfChannel#tile_padding}.
	// Experimental.
	TilePadding *string `field:"optional" json:"tilePadding" yaml:"tilePadding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#tile_width TfChannel#tile_width}.
	// Experimental.
	TileWidth *float64 `field:"optional" json:"tileWidth" yaml:"tileWidth"`
	// timecode_burnin_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timecode_burnin_settings TfChannel#timecode_burnin_settings}
	// Experimental.
	TimecodeBurninSettings *TfChannel_TimecodeBurninSettingsProperty `field:"optional" json:"timecodeBurninSettings" yaml:"timecodeBurninSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timecode_insertion TfChannel#timecode_insertion}.
	// Experimental.
	TimecodeInsertion *string `field:"optional" json:"timecodeInsertion" yaml:"timecodeInsertion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#treeblock_size TfChannel#treeblock_size}.
	// Experimental.
	TreeblockSize *string `field:"optional" json:"treeblockSize" yaml:"treeblockSize"`
}


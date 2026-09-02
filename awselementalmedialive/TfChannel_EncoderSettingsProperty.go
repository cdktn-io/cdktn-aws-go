package awselementalmedialive


// Experimental.
type TfChannel_EncoderSettingsProperty struct {
	// output_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_groups TfChannel#output_groups}
	// Experimental.
	OutputGroups interface{} `field:"required" json:"outputGroups" yaml:"outputGroups"`
	// timecode_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timecode_config TfChannel#timecode_config}
	// Experimental.
	TimecodeConfig *TfChannel_TimecodeConfigProperty `field:"required" json:"timecodeConfig" yaml:"timecodeConfig"`
	// audio_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_descriptions TfChannel#audio_descriptions}
	// Experimental.
	AudioDescriptions interface{} `field:"optional" json:"audioDescriptions" yaml:"audioDescriptions"`
	// avail_blanking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#avail_blanking TfChannel#avail_blanking}
	// Experimental.
	AvailBlanking *TfChannel_AvailBlankingProperty `field:"optional" json:"availBlanking" yaml:"availBlanking"`
	// caption_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_descriptions TfChannel#caption_descriptions}
	// Experimental.
	CaptionDescriptions interface{} `field:"optional" json:"captionDescriptions" yaml:"captionDescriptions"`
	// global_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#global_configuration TfChannel#global_configuration}
	// Experimental.
	GlobalConfiguration *TfChannel_GlobalConfigurationProperty `field:"optional" json:"globalConfiguration" yaml:"globalConfiguration"`
	// motion_graphics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#motion_graphics_configuration TfChannel#motion_graphics_configuration}
	// Experimental.
	MotionGraphicsConfiguration *TfChannel_MotionGraphicsConfigurationProperty `field:"optional" json:"motionGraphicsConfiguration" yaml:"motionGraphicsConfiguration"`
	// nielsen_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_configuration TfChannel#nielsen_configuration}
	// Experimental.
	NielsenConfiguration *TfChannel_NielsenConfigurationProperty `field:"optional" json:"nielsenConfiguration" yaml:"nielsenConfiguration"`
	// video_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_descriptions TfChannel#video_descriptions}
	// Experimental.
	VideoDescriptions interface{} `field:"optional" json:"videoDescriptions" yaml:"videoDescriptions"`
}


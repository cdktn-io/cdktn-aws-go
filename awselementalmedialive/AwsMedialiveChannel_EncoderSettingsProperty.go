package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_EncoderSettingsProperty struct {
	// output_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#output_groups AwsMedialiveChannel#output_groups}
	// Experimental.
	OutputGroups interface{} `field:"required" json:"outputGroups" yaml:"outputGroups"`
	// timecode_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#timecode_config AwsMedialiveChannel#timecode_config}
	// Experimental.
	TimecodeConfig *AwsMedialiveChannel_TimecodeConfigProperty `field:"required" json:"timecodeConfig" yaml:"timecodeConfig"`
	// audio_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#audio_descriptions AwsMedialiveChannel#audio_descriptions}
	// Experimental.
	AudioDescriptions interface{} `field:"optional" json:"audioDescriptions" yaml:"audioDescriptions"`
	// avail_blanking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#avail_blanking AwsMedialiveChannel#avail_blanking}
	// Experimental.
	AvailBlanking *AwsMedialiveChannel_AvailBlankingProperty `field:"optional" json:"availBlanking" yaml:"availBlanking"`
	// caption_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#caption_descriptions AwsMedialiveChannel#caption_descriptions}
	// Experimental.
	CaptionDescriptions interface{} `field:"optional" json:"captionDescriptions" yaml:"captionDescriptions"`
	// global_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#global_configuration AwsMedialiveChannel#global_configuration}
	// Experimental.
	GlobalConfiguration *AwsMedialiveChannel_GlobalConfigurationProperty `field:"optional" json:"globalConfiguration" yaml:"globalConfiguration"`
	// motion_graphics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#motion_graphics_configuration AwsMedialiveChannel#motion_graphics_configuration}
	// Experimental.
	MotionGraphicsConfiguration *AwsMedialiveChannel_MotionGraphicsConfigurationProperty `field:"optional" json:"motionGraphicsConfiguration" yaml:"motionGraphicsConfiguration"`
	// nielsen_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_configuration AwsMedialiveChannel#nielsen_configuration}
	// Experimental.
	NielsenConfiguration *AwsMedialiveChannel_NielsenConfigurationProperty `field:"optional" json:"nielsenConfiguration" yaml:"nielsenConfiguration"`
	// video_descriptions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_descriptions AwsMedialiveChannel#video_descriptions}
	// Experimental.
	VideoDescriptions interface{} `field:"optional" json:"videoDescriptions" yaml:"videoDescriptions"`
}


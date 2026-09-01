package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty struct {
	// frame_capture_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_settings AwsMedialiveChannel#frame_capture_settings}
	// Experimental.
	FrameCaptureSettings *AwsMedialiveChannel_FrameCaptureSettingsProperty `field:"optional" json:"frameCaptureSettings" yaml:"frameCaptureSettings"`
	// h264_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#h264_settings AwsMedialiveChannel#h264_settings}
	// Experimental.
	H264Settings *AwsMedialiveChannel_H264SettingsProperty `field:"optional" json:"h264Settings" yaml:"h264Settings"`
	// h265_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#h265_settings AwsMedialiveChannel#h265_settings}
	// Experimental.
	H265Settings *AwsMedialiveChannel_H265SettingsProperty `field:"optional" json:"h265Settings" yaml:"h265Settings"`
}


package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_FrameCaptureGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsMedialiveChannel#destination}
	// Experimental.
	Destination *AwsMedialiveChannel_EncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// frame_capture_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_cdn_settings AwsMedialiveChannel#frame_capture_cdn_settings}
	// Experimental.
	FrameCaptureCdnSettings *AwsMedialiveChannel_FrameCaptureCdnSettingsProperty `field:"optional" json:"frameCaptureCdnSettings" yaml:"frameCaptureCdnSettings"`
}


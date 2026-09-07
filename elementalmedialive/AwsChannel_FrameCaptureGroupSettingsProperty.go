package elementalmedialive


// Experimental.
type AwsChannel_FrameCaptureGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination AwsChannel#destination}
	// Experimental.
	Destination *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// frame_capture_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_cdn_settings AwsChannel#frame_capture_cdn_settings}
	// Experimental.
	FrameCaptureCdnSettings *AwsChannel_FrameCaptureCdnSettingsProperty `field:"optional" json:"frameCaptureCdnSettings" yaml:"frameCaptureCdnSettings"`
}


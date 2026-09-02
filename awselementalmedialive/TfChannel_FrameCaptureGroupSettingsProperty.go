package awselementalmedialive


// Experimental.
type TfChannel_FrameCaptureGroupSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination TfChannel#destination}
	// Experimental.
	Destination *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsFrameCaptureGroupSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// frame_capture_cdn_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_cdn_settings TfChannel#frame_capture_cdn_settings}
	// Experimental.
	FrameCaptureCdnSettings *TfChannel_FrameCaptureCdnSettingsProperty `field:"optional" json:"frameCaptureCdnSettings" yaml:"frameCaptureCdnSettings"`
}


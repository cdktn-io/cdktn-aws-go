package awselementalmedialive


// Experimental.
type TfChannel_FrameCaptureCdnSettingsProperty struct {
	// frame_capture_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#frame_capture_s3_settings TfChannel#frame_capture_s3_settings}
	// Experimental.
	FrameCaptureS3Settings *TfChannel_FrameCaptureS3SettingsProperty `field:"optional" json:"frameCaptureS3Settings" yaml:"frameCaptureS3Settings"`
}


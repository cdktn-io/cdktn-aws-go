package elementalmedialive


// Experimental.
type AwsChannel_VideoBlackSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#black_detect_threshold AwsChannel#black_detect_threshold}.
	// Experimental.
	BlackDetectThreshold *float64 `field:"optional" json:"blackDetectThreshold" yaml:"blackDetectThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#video_black_threshold_msec AwsChannel#video_black_threshold_msec}.
	// Experimental.
	VideoBlackThresholdMsec *float64 `field:"optional" json:"videoBlackThresholdMsec" yaml:"videoBlackThresholdMsec"`
}


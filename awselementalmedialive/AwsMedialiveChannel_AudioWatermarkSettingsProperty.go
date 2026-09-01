package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_AudioWatermarkSettingsProperty struct {
	// nielsen_watermarks_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_watermarks_settings AwsMedialiveChannel#nielsen_watermarks_settings}
	// Experimental.
	NielsenWatermarksSettings *AwsMedialiveChannel_NielsenWatermarksSettingsProperty `field:"optional" json:"nielsenWatermarksSettings" yaml:"nielsenWatermarksSettings"`
}


package elementalmedialive


// Experimental.
type AwsChannel_AudioWatermarkSettingsProperty struct {
	// nielsen_watermarks_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_watermarks_settings AwsChannel#nielsen_watermarks_settings}
	// Experimental.
	NielsenWatermarksSettings *AwsChannel_NielsenWatermarksSettingsProperty `field:"optional" json:"nielsenWatermarksSettings" yaml:"nielsenWatermarksSettings"`
}


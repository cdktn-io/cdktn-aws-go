package awselementalmedialive


// Experimental.
type TfChannel_AudioWatermarkSettingsProperty struct {
	// nielsen_watermarks_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#nielsen_watermarks_settings TfChannel#nielsen_watermarks_settings}
	// Experimental.
	NielsenWatermarksSettings *TfChannel_NielsenWatermarksSettingsProperty `field:"optional" json:"nielsenWatermarksSettings" yaml:"nielsenWatermarksSettings"`
}


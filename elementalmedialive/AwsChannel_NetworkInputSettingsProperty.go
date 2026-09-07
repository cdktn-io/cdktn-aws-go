package elementalmedialive


// Experimental.
type AwsChannel_NetworkInputSettingsProperty struct {
	// hls_input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_input_settings AwsChannel#hls_input_settings}
	// Experimental.
	HlsInputSettings *AwsChannel_HlsInputSettingsProperty `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#server_validation AwsChannel#server_validation}.
	// Experimental.
	ServerValidation *string `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}


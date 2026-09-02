package awselementalmedialive


// Experimental.
type TfChannel_NetworkInputSettingsProperty struct {
	// hls_input_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#hls_input_settings TfChannel#hls_input_settings}
	// Experimental.
	HlsInputSettings *TfChannel_HlsInputSettingsProperty `field:"optional" json:"hlsInputSettings" yaml:"hlsInputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#server_validation TfChannel#server_validation}.
	// Experimental.
	ServerValidation *string `field:"optional" json:"serverValidation" yaml:"serverValidation"`
}


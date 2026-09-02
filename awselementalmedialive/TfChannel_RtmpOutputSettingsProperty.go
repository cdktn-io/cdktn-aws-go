package awselementalmedialive


// Experimental.
type TfChannel_RtmpOutputSettingsProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#destination TfChannel#destination}
	// Experimental.
	Destination *TfChannel_EncoderSettingsOutputGroupsOutputsOutputSettingsRtmpOutputSettingsDestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#certificate_mode TfChannel#certificate_mode}.
	// Experimental.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#connection_retry_interval TfChannel#connection_retry_interval}.
	// Experimental.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#num_retries TfChannel#num_retries}.
	// Experimental.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
}


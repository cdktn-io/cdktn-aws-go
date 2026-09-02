package awselementalmedialive


// Experimental.
type TfChannel_HlsAkamaiSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#connection_retry_interval TfChannel#connection_retry_interval}.
	// Experimental.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#filecache_duration TfChannel#filecache_duration}.
	// Experimental.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#http_transfer_mode TfChannel#http_transfer_mode}.
	// Experimental.
	HttpTransferMode *string `field:"optional" json:"httpTransferMode" yaml:"httpTransferMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#num_retries TfChannel#num_retries}.
	// Experimental.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#restart_delay TfChannel#restart_delay}.
	// Experimental.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#salt TfChannel#salt}.
	// Experimental.
	Salt *string `field:"optional" json:"salt" yaml:"salt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#token TfChannel#token}.
	// Experimental.
	Token *string `field:"optional" json:"token" yaml:"token"`
}


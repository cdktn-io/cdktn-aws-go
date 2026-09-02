package awselementalmedialive


// Experimental.
type TfChannel_StaticKeySettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#static_key_value TfChannel#static_key_value}.
	// Experimental.
	StaticKeyValue *string `field:"required" json:"staticKeyValue" yaml:"staticKeyValue"`
	// key_provider_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#key_provider_server TfChannel#key_provider_server}
	// Experimental.
	KeyProviderServer *TfChannel_KeyProviderServerProperty `field:"optional" json:"keyProviderServer" yaml:"keyProviderServer"`
}


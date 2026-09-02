package awschime


// Experimental.
type TfVoiceConnectorTerminationCredentials_CredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination_credentials#password TfVoiceConnectorTerminationCredentials#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination_credentials#username TfVoiceConnectorTerminationCredentials#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}


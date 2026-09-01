package awschime


// Experimental.
type AwsChimeVoiceConnectorTerminationCredentials_CredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination_credentials#password AwsChimeVoiceConnectorTerminationCredentials#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_termination_credentials#username AwsChimeVoiceConnectorTerminationCredentials#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}


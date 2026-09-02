package awschime


// Experimental.
type TfVoiceConnectorGroup_ConnectorProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_group#priority TfVoiceConnectorGroup#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_group#voice_connector_id TfVoiceConnectorGroup#voice_connector_id}.
	// Experimental.
	VoiceConnectorId *string `field:"required" json:"voiceConnectorId" yaml:"voiceConnectorId"`
}


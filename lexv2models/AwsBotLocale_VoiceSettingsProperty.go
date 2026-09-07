package lexv2models


// Experimental.
type AwsBotLocale_VoiceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_bot_locale#voice_id AwsBotLocale#voice_id}.
	// Experimental.
	VoiceId *string `field:"required" json:"voiceId" yaml:"voiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_bot_locale#engine AwsBotLocale#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
}


package awslexv2models


// Experimental.
type TfBotLocale_VoiceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_bot_locale#voice_id TfBotLocale#voice_id}.
	// Experimental.
	VoiceId *string `field:"required" json:"voiceId" yaml:"voiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_bot_locale#engine TfBotLocale#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
}


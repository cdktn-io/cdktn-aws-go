package awslexv2models


// Experimental.
type AwsLexv2ModelsBotLocale_VoiceSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_bot_locale#voice_id AwsLexv2ModelsBotLocale#voice_id}.
	// Experimental.
	VoiceId *string `field:"required" json:"voiceId" yaml:"voiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_bot_locale#engine AwsLexv2ModelsBotLocale#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
}


package awslexv2models


// Experimental.
type TfIntent_AllowedInputTypesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#allow_audio_input TfIntent#allow_audio_input}.
	// Experimental.
	AllowAudioInput interface{} `field:"required" json:"allowAudioInput" yaml:"allowAudioInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#allow_dtmf_input TfIntent#allow_dtmf_input}.
	// Experimental.
	AllowDtmfInput interface{} `field:"required" json:"allowDtmfInput" yaml:"allowDtmfInput"`
}


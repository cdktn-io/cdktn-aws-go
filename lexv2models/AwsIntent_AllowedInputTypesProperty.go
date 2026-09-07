package lexv2models


// Experimental.
type AwsIntent_AllowedInputTypesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#allow_audio_input AwsIntent#allow_audio_input}.
	// Experimental.
	AllowAudioInput interface{} `field:"required" json:"allowAudioInput" yaml:"allowAudioInput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#allow_dtmf_input AwsIntent#allow_dtmf_input}.
	// Experimental.
	AllowDtmfInput interface{} `field:"required" json:"allowDtmfInput" yaml:"allowDtmfInput"`
}


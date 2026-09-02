package awschime


// Experimental.
type TfVoiceConnectorOrigination_RouteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_origination#host TfVoiceConnectorOrigination#host}.
	// Experimental.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_origination#priority TfVoiceConnectorOrigination#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_origination#protocol TfVoiceConnectorOrigination#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_origination#weight TfVoiceConnectorOrigination#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chime_voice_connector_origination#port TfVoiceConnectorOrigination#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}


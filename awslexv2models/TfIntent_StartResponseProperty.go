package awslexv2models


// Experimental.
type TfIntent_StartResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#allow_interrupt TfIntent#allow_interrupt}.
	// Experimental.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#delay_in_seconds TfIntent#delay_in_seconds}.
	// Experimental.
	DelayInSeconds *float64 `field:"optional" json:"delayInSeconds" yaml:"delayInSeconds"`
	// message_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#message_group TfIntent#message_group}
	// Experimental.
	MessageGroup interface{} `field:"optional" json:"messageGroup" yaml:"messageGroup"`
}


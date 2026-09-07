package lexv2models


// Experimental.
type AwsIntent_StartResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#allow_interrupt AwsIntent#allow_interrupt}.
	// Experimental.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#delay_in_seconds AwsIntent#delay_in_seconds}.
	// Experimental.
	DelayInSeconds *float64 `field:"optional" json:"delayInSeconds" yaml:"delayInSeconds"`
	// message_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#message_group AwsIntent#message_group}
	// Experimental.
	MessageGroup interface{} `field:"optional" json:"messageGroup" yaml:"messageGroup"`
}


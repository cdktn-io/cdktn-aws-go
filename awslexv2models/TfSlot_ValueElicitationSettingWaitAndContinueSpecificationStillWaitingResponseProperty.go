package awslexv2models


// Experimental.
type TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#frequency_in_seconds TfSlot#frequency_in_seconds}.
	// Experimental.
	FrequencyInSeconds *float64 `field:"required" json:"frequencyInSeconds" yaml:"frequencyInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#timeout_in_seconds TfSlot#timeout_in_seconds}.
	// Experimental.
	TimeoutInSeconds *float64 `field:"required" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#allow_interrupt TfSlot#allow_interrupt}.
	// Experimental.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// message_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_slot#message_group TfSlot#message_group}
	// Experimental.
	MessageGroup interface{} `field:"optional" json:"messageGroup" yaml:"messageGroup"`
}


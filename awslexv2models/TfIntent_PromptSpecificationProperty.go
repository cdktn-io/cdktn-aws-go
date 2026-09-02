package awslexv2models


// Experimental.
type TfIntent_PromptSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#max_retries TfIntent#max_retries}.
	// Experimental.
	MaxRetries *float64 `field:"required" json:"maxRetries" yaml:"maxRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#allow_interrupt TfIntent#allow_interrupt}.
	// Experimental.
	AllowInterrupt interface{} `field:"optional" json:"allowInterrupt" yaml:"allowInterrupt"`
	// message_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#message_group TfIntent#message_group}
	// Experimental.
	MessageGroup interface{} `field:"optional" json:"messageGroup" yaml:"messageGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#message_selection_strategy TfIntent#message_selection_strategy}.
	// Experimental.
	MessageSelectionStrategy *string `field:"optional" json:"messageSelectionStrategy" yaml:"messageSelectionStrategy"`
	// prompt_attempts_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#prompt_attempts_specification TfIntent#prompt_attempts_specification}
	// Experimental.
	PromptAttemptsSpecification interface{} `field:"optional" json:"promptAttemptsSpecification" yaml:"promptAttemptsSpecification"`
}


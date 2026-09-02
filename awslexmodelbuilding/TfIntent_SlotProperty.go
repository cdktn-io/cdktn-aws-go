package awslexmodelbuilding


// Experimental.
type TfIntent_SlotProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#name TfIntent#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#slot_constraint TfIntent#slot_constraint}.
	// Experimental.
	SlotConstraint *string `field:"required" json:"slotConstraint" yaml:"slotConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#slot_type TfIntent#slot_type}.
	// Experimental.
	SlotType *string `field:"required" json:"slotType" yaml:"slotType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#description TfIntent#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#priority TfIntent#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#response_card TfIntent#response_card}.
	// Experimental.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#sample_utterances TfIntent#sample_utterances}.
	// Experimental.
	SampleUtterances *[]*string `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#slot_type_version TfIntent#slot_type_version}.
	// Experimental.
	SlotTypeVersion *string `field:"optional" json:"slotTypeVersion" yaml:"slotTypeVersion"`
	// value_elicitation_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#value_elicitation_prompt TfIntent#value_elicitation_prompt}
	// Experimental.
	ValueElicitationPrompt *TfIntent_ValueElicitationPromptProperty `field:"optional" json:"valueElicitationPrompt" yaml:"valueElicitationPrompt"`
}


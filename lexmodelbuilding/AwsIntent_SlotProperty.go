package lexmodelbuilding


// Experimental.
type AwsIntent_SlotProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#name AwsIntent#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#slot_constraint AwsIntent#slot_constraint}.
	// Experimental.
	SlotConstraint *string `field:"required" json:"slotConstraint" yaml:"slotConstraint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#slot_type AwsIntent#slot_type}.
	// Experimental.
	SlotType *string `field:"required" json:"slotType" yaml:"slotType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#description AwsIntent#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#priority AwsIntent#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#response_card AwsIntent#response_card}.
	// Experimental.
	ResponseCard *string `field:"optional" json:"responseCard" yaml:"responseCard"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#sample_utterances AwsIntent#sample_utterances}.
	// Experimental.
	SampleUtterances *[]*string `field:"optional" json:"sampleUtterances" yaml:"sampleUtterances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#slot_type_version AwsIntent#slot_type_version}.
	// Experimental.
	SlotTypeVersion *string `field:"optional" json:"slotTypeVersion" yaml:"slotTypeVersion"`
	// value_elicitation_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent#value_elicitation_prompt AwsIntent#value_elicitation_prompt}
	// Experimental.
	ValueElicitationPrompt *AwsIntent_ValueElicitationPromptProperty `field:"optional" json:"valueElicitationPrompt" yaml:"valueElicitationPrompt"`
}


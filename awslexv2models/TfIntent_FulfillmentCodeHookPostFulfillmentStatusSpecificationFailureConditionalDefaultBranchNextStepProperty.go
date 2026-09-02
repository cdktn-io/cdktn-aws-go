package awslexv2models


// Experimental.
type TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalDefaultBranchNextStepProperty struct {
	// dialog_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#dialog_action TfIntent#dialog_action}
	// Experimental.
	DialogAction interface{} `field:"optional" json:"dialogAction" yaml:"dialogAction"`
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#intent TfIntent#intent}
	// Experimental.
	Intent interface{} `field:"optional" json:"intent" yaml:"intent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#session_attributes TfIntent#session_attributes}.
	// Experimental.
	SessionAttributes *map[string]*string `field:"optional" json:"sessionAttributes" yaml:"sessionAttributes"`
}


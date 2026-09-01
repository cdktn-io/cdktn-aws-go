package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepProperty struct {
	// dialog_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#dialog_action AwsLexv2ModelsIntent#dialog_action}
	// Experimental.
	DialogAction interface{} `field:"optional" json:"dialogAction" yaml:"dialogAction"`
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#intent AwsLexv2ModelsIntent#intent}
	// Experimental.
	Intent interface{} `field:"optional" json:"intent" yaml:"intent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#session_attributes AwsLexv2ModelsIntent#session_attributes}.
	// Experimental.
	SessionAttributes *map[string]*string `field:"optional" json:"sessionAttributes" yaml:"sessionAttributes"`
}


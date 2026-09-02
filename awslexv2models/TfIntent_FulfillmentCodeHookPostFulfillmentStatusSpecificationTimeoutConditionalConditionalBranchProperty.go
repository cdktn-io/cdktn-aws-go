package awslexv2models


// Experimental.
type TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#name TfIntent#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#condition TfIntent#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// next_step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#next_step TfIntent#next_step}
	// Experimental.
	NextStep interface{} `field:"optional" json:"nextStep" yaml:"nextStep"`
	// response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#response TfIntent#response}
	// Experimental.
	Response interface{} `field:"optional" json:"response" yaml:"response"`
}


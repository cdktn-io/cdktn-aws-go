package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessConditionalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active AwsLexv2ModelsIntent#active}.
	// Experimental.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// conditional_branch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#conditional_branch AwsLexv2ModelsIntent#conditional_branch}
	// Experimental.
	ConditionalBranch interface{} `field:"optional" json:"conditionalBranch" yaml:"conditionalBranch"`
	// default_branch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#default_branch AwsLexv2ModelsIntent#default_branch}
	// Experimental.
	DefaultBranch interface{} `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
}


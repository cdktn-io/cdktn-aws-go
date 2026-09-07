package lexv2models


// Experimental.
type AwsIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active AwsIntent#active}.
	// Experimental.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// conditional_branch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#conditional_branch AwsIntent#conditional_branch}
	// Experimental.
	ConditionalBranch interface{} `field:"optional" json:"conditionalBranch" yaml:"conditionalBranch"`
	// default_branch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#default_branch AwsIntent#default_branch}
	// Experimental.
	DefaultBranch interface{} `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
}


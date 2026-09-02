package awsecr


// Experimental.
type DataTfLifecyclePolicyDocument_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#priority DataTfLifecyclePolicyDocument#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#action DataTfLifecyclePolicyDocument#action}
	// Experimental.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#description DataTfLifecyclePolicyDocument#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecr_lifecycle_policy_document#selection DataTfLifecyclePolicyDocument#selection}
	// Experimental.
	Selection interface{} `field:"optional" json:"selection" yaml:"selection"`
}


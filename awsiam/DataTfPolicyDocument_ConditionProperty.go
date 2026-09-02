package awsiam


// Experimental.
type DataTfPolicyDocument_ConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#test DataTfPolicyDocument#test}.
	// Experimental.
	Test *string `field:"required" json:"test" yaml:"test"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#values DataTfPolicyDocument#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#variable DataTfPolicyDocument#variable}.
	// Experimental.
	Variable *string `field:"required" json:"variable" yaml:"variable"`
}


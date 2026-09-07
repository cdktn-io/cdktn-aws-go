package iam


// Experimental.
type DataAwsPolicyDocument_ConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#test DataAwsPolicyDocument#test}.
	// Experimental.
	Test *string `field:"required" json:"test" yaml:"test"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#values DataAwsPolicyDocument#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#variable DataAwsPolicyDocument#variable}.
	// Experimental.
	Variable *string `field:"required" json:"variable" yaml:"variable"`
}


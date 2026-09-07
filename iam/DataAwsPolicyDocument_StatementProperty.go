package iam


// Experimental.
type DataAwsPolicyDocument_StatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#actions DataAwsPolicyDocument#actions}.
	// Experimental.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#condition DataAwsPolicyDocument#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#effect DataAwsPolicyDocument#effect}.
	// Experimental.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#not_actions DataAwsPolicyDocument#not_actions}.
	// Experimental.
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// not_principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#not_principals DataAwsPolicyDocument#not_principals}
	// Experimental.
	NotPrincipals interface{} `field:"optional" json:"notPrincipals" yaml:"notPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#not_resources DataAwsPolicyDocument#not_resources}.
	// Experimental.
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#principals DataAwsPolicyDocument#principals}
	// Experimental.
	Principals interface{} `field:"optional" json:"principals" yaml:"principals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#resources DataAwsPolicyDocument#resources}.
	// Experimental.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#sid DataAwsPolicyDocument#sid}.
	// Experimental.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}


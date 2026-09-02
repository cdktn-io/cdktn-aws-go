package awsiam


// Experimental.
type DataTfPolicyDocument_StatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#actions DataTfPolicyDocument#actions}.
	// Experimental.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#condition DataTfPolicyDocument#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#effect DataTfPolicyDocument#effect}.
	// Experimental.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#not_actions DataTfPolicyDocument#not_actions}.
	// Experimental.
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// not_principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#not_principals DataTfPolicyDocument#not_principals}
	// Experimental.
	NotPrincipals interface{} `field:"optional" json:"notPrincipals" yaml:"notPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#not_resources DataTfPolicyDocument#not_resources}.
	// Experimental.
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#principals DataTfPolicyDocument#principals}
	// Experimental.
	Principals interface{} `field:"optional" json:"principals" yaml:"principals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#resources DataTfPolicyDocument#resources}.
	// Experimental.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#sid DataTfPolicyDocument#sid}.
	// Experimental.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}


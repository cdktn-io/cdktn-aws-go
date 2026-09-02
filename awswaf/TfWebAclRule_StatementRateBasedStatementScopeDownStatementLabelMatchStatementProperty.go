package awswaf


// Experimental.
type TfWebAclRule_StatementRateBasedStatementScopeDownStatementLabelMatchStatementProperty struct {
	// String to match against. Must be 1-1024 characters and match pattern ^[0-9A-Za-z_\-:]+$.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#key TfWebAclRule#key}
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Specify whether to match using the label name or just the namespace. Valid values: LABEL, NAMESPACE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#scope TfWebAclRule#scope}
	// Experimental.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}


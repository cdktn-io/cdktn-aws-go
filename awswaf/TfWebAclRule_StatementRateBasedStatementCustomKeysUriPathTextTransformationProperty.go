package awswaf


// Experimental.
type TfWebAclRule_StatementRateBasedStatementCustomKeysUriPathTextTransformationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#priority TfWebAclRule#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#type TfWebAclRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}


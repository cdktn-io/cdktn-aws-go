package awswaf


// Experimental.
type TfWebAclRule_StatementRateBasedStatementForwardedIpConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#fallback_behavior TfWebAclRule#fallback_behavior}.
	// Experimental.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#header_name TfWebAclRule#header_name}.
	// Experimental.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}


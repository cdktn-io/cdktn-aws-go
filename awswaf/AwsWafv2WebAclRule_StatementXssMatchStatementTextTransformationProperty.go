package awswaf


// Experimental.
type AwsWafv2WebAclRule_StatementXssMatchStatementTextTransformationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#priority AwsWafv2WebAclRule#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#type AwsWafv2WebAclRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}


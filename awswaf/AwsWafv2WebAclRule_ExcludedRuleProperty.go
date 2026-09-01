package awswaf


// Experimental.
type AwsWafv2WebAclRule_ExcludedRuleProperty struct {
	// Name of the rule to exclude (1-128 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#name AwsWafv2WebAclRule#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}


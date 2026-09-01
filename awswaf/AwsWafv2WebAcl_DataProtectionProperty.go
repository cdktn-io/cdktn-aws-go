package awswaf


// Experimental.
type AwsWafv2WebAcl_DataProtectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#action AwsWafv2WebAcl#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#field AwsWafv2WebAcl#field}
	// Experimental.
	Field *AwsWafv2WebAcl_FieldProperty `field:"required" json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#exclude_rate_based_details AwsWafv2WebAcl#exclude_rate_based_details}.
	// Experimental.
	ExcludeRateBasedDetails interface{} `field:"optional" json:"excludeRateBasedDetails" yaml:"excludeRateBasedDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#exclude_rule_match_details AwsWafv2WebAcl#exclude_rule_match_details}.
	// Experimental.
	ExcludeRuleMatchDetails interface{} `field:"optional" json:"excludeRuleMatchDetails" yaml:"excludeRuleMatchDetails"`
}


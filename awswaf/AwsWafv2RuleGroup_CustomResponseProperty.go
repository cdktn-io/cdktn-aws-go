package awswaf


// Experimental.
type AwsWafv2RuleGroup_CustomResponseProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#response_code AwsWafv2RuleGroup#response_code}.
	// Experimental.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#custom_response_body_key AwsWafv2RuleGroup#custom_response_body_key}.
	// Experimental.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#response_header AwsWafv2RuleGroup#response_header}
	// Experimental.
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
}


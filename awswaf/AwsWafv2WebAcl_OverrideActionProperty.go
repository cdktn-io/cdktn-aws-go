package awswaf


// Experimental.
type AwsWafv2WebAcl_OverrideActionProperty struct {
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#count AwsWafv2WebAcl#count}
	// Experimental.
	Count *AwsWafv2WebAcl_RuleOverrideActionCountProperty `field:"optional" json:"count" yaml:"count"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#none AwsWafv2WebAcl#none}
	// Experimental.
	None *AwsWafv2WebAcl_NoneProperty `field:"optional" json:"none" yaml:"none"`
}


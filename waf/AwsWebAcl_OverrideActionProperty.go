package waf


// Experimental.
type AwsWebAcl_OverrideActionProperty struct {
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#count AwsWebAcl#count}
	// Experimental.
	Count *AwsWebAcl_RuleOverrideActionCountProperty `field:"optional" json:"count" yaml:"count"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#none AwsWebAcl#none}
	// Experimental.
	None *AwsWebAcl_NoneProperty `field:"optional" json:"none" yaml:"none"`
}


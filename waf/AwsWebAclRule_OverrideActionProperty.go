package waf


// Experimental.
type AwsWebAclRule_OverrideActionProperty struct {
	// count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#count AwsWebAclRule#count}
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// none block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#none AwsWebAclRule#none}
	// Experimental.
	None interface{} `field:"optional" json:"none" yaml:"none"`
}


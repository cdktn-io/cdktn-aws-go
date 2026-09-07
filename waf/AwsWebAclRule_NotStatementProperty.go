package waf


// Experimental.
type AwsWebAclRule_NotStatementProperty struct {
	// statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#statement AwsWebAclRule#statement}
	// Experimental.
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}


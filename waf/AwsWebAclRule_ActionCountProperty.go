package waf


// Experimental.
type AwsWebAclRule_ActionCountProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#custom_request_handling AwsWebAclRule#custom_request_handling}
	// Experimental.
	CustomRequestHandling interface{} `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}


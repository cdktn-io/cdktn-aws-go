package waf


// Experimental.
type AwsWebAcl_RuleActionAllowProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#custom_request_handling AwsWebAcl#custom_request_handling}
	// Experimental.
	CustomRequestHandling *AwsWebAcl_RuleActionAllowCustomRequestHandlingProperty `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}


package awswaf


// Experimental.
type AwsWafv2WebAcl_CaptchaProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#custom_request_handling AwsWafv2WebAcl#custom_request_handling}
	// Experimental.
	CustomRequestHandling *AwsWafv2WebAcl_RuleActionCaptchaCustomRequestHandlingProperty `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}


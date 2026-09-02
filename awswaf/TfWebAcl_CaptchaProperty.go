package awswaf


// Experimental.
type TfWebAcl_CaptchaProperty struct {
	// custom_request_handling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#custom_request_handling TfWebAcl#custom_request_handling}
	// Experimental.
	CustomRequestHandling *TfWebAcl_RuleActionCaptchaCustomRequestHandlingProperty `field:"optional" json:"customRequestHandling" yaml:"customRequestHandling"`
}


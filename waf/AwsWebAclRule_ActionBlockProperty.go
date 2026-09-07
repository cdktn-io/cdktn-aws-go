package waf


// Experimental.
type AwsWebAclRule_ActionBlockProperty struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#custom_response AwsWebAclRule#custom_response}
	// Experimental.
	CustomResponse interface{} `field:"optional" json:"customResponse" yaml:"customResponse"`
}


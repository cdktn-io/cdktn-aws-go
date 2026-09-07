package waf


// Experimental.
type AwsWebAcl_DefaultActionBlockProperty struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#custom_response AwsWebAcl#custom_response}
	// Experimental.
	CustomResponse *AwsWebAcl_DefaultActionBlockCustomResponseProperty `field:"optional" json:"customResponse" yaml:"customResponse"`
}


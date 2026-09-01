package awswaf


// Experimental.
type AwsWafv2WebAcl_RuleActionBlockProperty struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl#custom_response AwsWafv2WebAcl#custom_response}
	// Experimental.
	CustomResponse *AwsWafv2WebAcl_RuleActionBlockCustomResponseProperty `field:"optional" json:"customResponse" yaml:"customResponse"`
}


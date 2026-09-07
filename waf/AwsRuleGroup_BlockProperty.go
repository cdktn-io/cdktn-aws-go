package waf


// Experimental.
type AwsRuleGroup_BlockProperty struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#custom_response AwsRuleGroup#custom_response}
	// Experimental.
	CustomResponse *AwsRuleGroup_CustomResponseProperty `field:"optional" json:"customResponse" yaml:"customResponse"`
}


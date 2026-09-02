package awswaf


// Experimental.
type TfRuleGroup_BlockProperty struct {
	// custom_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#custom_response TfRuleGroup#custom_response}
	// Experimental.
	CustomResponse *TfRuleGroup_CustomResponseProperty `field:"optional" json:"customResponse" yaml:"customResponse"`
}


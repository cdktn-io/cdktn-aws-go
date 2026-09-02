package awswaf


// Experimental.
type TfRuleGroup_CustomResponseBodyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#content TfRuleGroup#content}.
	// Experimental.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#content_type TfRuleGroup#content_type}.
	// Experimental.
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_rule_group#key TfRuleGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}


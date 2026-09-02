package awssesmailmanager


// Experimental.
type TfRuleSet_RuleConditionBooleanExpressionEvaluateIsInAddressListProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#address_lists TfRuleSet#address_lists}.
	// Experimental.
	AddressLists *[]*string `field:"required" json:"addressLists" yaml:"addressLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#attribute TfRuleSet#attribute}.
	// Experimental.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}


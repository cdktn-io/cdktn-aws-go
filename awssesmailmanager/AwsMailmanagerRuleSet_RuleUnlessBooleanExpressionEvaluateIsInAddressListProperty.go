package awssesmailmanager


// Experimental.
type AwsMailmanagerRuleSet_RuleUnlessBooleanExpressionEvaluateIsInAddressListProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#address_lists AwsMailmanagerRuleSet#address_lists}.
	// Experimental.
	AddressLists *[]*string `field:"required" json:"addressLists" yaml:"addressLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#attribute AwsMailmanagerRuleSet#attribute}.
	// Experimental.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
}


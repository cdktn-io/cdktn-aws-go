package sesmailmanager


// Experimental.
type AwsRuleSet_RuleProperty struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action AwsRuleSet#action}
	// Experimental.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#condition AwsRuleSet#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#name AwsRuleSet#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// unless block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#unless AwsRuleSet#unless}
	// Experimental.
	Unless interface{} `field:"optional" json:"unless" yaml:"unless"`
}


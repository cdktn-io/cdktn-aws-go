package awssesmailmanager


// Experimental.
type TfRuleSet_UnlessProperty struct {
	// boolean_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#boolean_expression TfRuleSet#boolean_expression}
	// Experimental.
	BooleanExpression interface{} `field:"optional" json:"booleanExpression" yaml:"booleanExpression"`
	// dmarc_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#dmarc_expression TfRuleSet#dmarc_expression}
	// Experimental.
	DmarcExpression interface{} `field:"optional" json:"dmarcExpression" yaml:"dmarcExpression"`
	// ip_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#ip_expression TfRuleSet#ip_expression}
	// Experimental.
	IpExpression interface{} `field:"optional" json:"ipExpression" yaml:"ipExpression"`
	// number_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#number_expression TfRuleSet#number_expression}
	// Experimental.
	NumberExpression interface{} `field:"optional" json:"numberExpression" yaml:"numberExpression"`
	// string_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#string_expression TfRuleSet#string_expression}
	// Experimental.
	StringExpression interface{} `field:"optional" json:"stringExpression" yaml:"stringExpression"`
	// verdict_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#verdict_expression TfRuleSet#verdict_expression}
	// Experimental.
	VerdictExpression interface{} `field:"optional" json:"verdictExpression" yaml:"verdictExpression"`
}


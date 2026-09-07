package networkfirewall


// Experimental.
type AwsRuleGroup_StatelessRulesAndCustomActionsProperty struct {
	// stateless_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#stateless_rule AwsRuleGroup#stateless_rule}
	// Experimental.
	StatelessRule interface{} `field:"required" json:"statelessRule" yaml:"statelessRule"`
	// custom_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#custom_action AwsRuleGroup#custom_action}
	// Experimental.
	CustomAction interface{} `field:"optional" json:"customAction" yaml:"customAction"`
}


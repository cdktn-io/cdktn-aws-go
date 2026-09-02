package awsnetworkfirewall


// Experimental.
type TfRuleGroup_RulesSourceListProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#generated_rules_type TfRuleGroup#generated_rules_type}.
	// Experimental.
	GeneratedRulesType *string `field:"required" json:"generatedRulesType" yaml:"generatedRulesType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#targets TfRuleGroup#targets}.
	// Experimental.
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#target_types TfRuleGroup#target_types}.
	// Experimental.
	TargetTypes *[]*string `field:"required" json:"targetTypes" yaml:"targetTypes"`
}


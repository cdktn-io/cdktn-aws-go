package sesmailmanager


// Experimental.
type AwsRuleSet_RelayProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#relay AwsRuleSet#relay}.
	// Experimental.
	Relay *string `field:"required" json:"relay" yaml:"relay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy AwsRuleSet#action_failure_policy}.
	// Experimental.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#mail_from AwsRuleSet#mail_from}.
	// Experimental.
	MailFrom *string `field:"optional" json:"mailFrom" yaml:"mailFrom"`
}


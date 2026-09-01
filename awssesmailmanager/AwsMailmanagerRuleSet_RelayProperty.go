package awssesmailmanager


// Experimental.
type AwsMailmanagerRuleSet_RelayProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#relay AwsMailmanagerRuleSet#relay}.
	// Experimental.
	Relay *string `field:"required" json:"relay" yaml:"relay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy AwsMailmanagerRuleSet#action_failure_policy}.
	// Experimental.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#mail_from AwsMailmanagerRuleSet#mail_from}.
	// Experimental.
	MailFrom *string `field:"optional" json:"mailFrom" yaml:"mailFrom"`
}


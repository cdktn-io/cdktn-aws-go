package networkfirewall


// Experimental.
type AwsRuleGroup_CustomActionProperty struct {
	// action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#action_definition AwsRuleGroup#action_definition}
	// Experimental.
	ActionDefinition *AwsRuleGroup_ActionDefinitionProperty `field:"required" json:"actionDefinition" yaml:"actionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#action_name AwsRuleGroup#action_name}.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
}


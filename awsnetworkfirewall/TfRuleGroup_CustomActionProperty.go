package awsnetworkfirewall


// Experimental.
type TfRuleGroup_CustomActionProperty struct {
	// action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#action_definition TfRuleGroup#action_definition}
	// Experimental.
	ActionDefinition *TfRuleGroup_ActionDefinitionProperty `field:"required" json:"actionDefinition" yaml:"actionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkfirewall_rule_group#action_name TfRuleGroup#action_name}.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
}


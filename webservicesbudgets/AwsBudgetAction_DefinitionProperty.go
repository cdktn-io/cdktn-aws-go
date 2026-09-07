package webservicesbudgets


// Experimental.
type AwsBudgetAction_DefinitionProperty struct {
	// iam_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#iam_action_definition AwsBudgetAction#iam_action_definition}
	// Experimental.
	IamActionDefinition *AwsBudgetAction_IamActionDefinitionProperty `field:"optional" json:"iamActionDefinition" yaml:"iamActionDefinition"`
	// scp_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#scp_action_definition AwsBudgetAction#scp_action_definition}
	// Experimental.
	ScpActionDefinition *AwsBudgetAction_ScpActionDefinitionProperty `field:"optional" json:"scpActionDefinition" yaml:"scpActionDefinition"`
	// ssm_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#ssm_action_definition AwsBudgetAction#ssm_action_definition}
	// Experimental.
	SsmActionDefinition *AwsBudgetAction_SsmActionDefinitionProperty `field:"optional" json:"ssmActionDefinition" yaml:"ssmActionDefinition"`
}


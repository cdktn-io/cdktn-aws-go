package awswebservicesbudgets


// Experimental.
type TfBudgetAction_DefinitionProperty struct {
	// iam_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#iam_action_definition TfBudgetAction#iam_action_definition}
	// Experimental.
	IamActionDefinition *TfBudgetAction_IamActionDefinitionProperty `field:"optional" json:"iamActionDefinition" yaml:"iamActionDefinition"`
	// scp_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#scp_action_definition TfBudgetAction#scp_action_definition}
	// Experimental.
	ScpActionDefinition *TfBudgetAction_ScpActionDefinitionProperty `field:"optional" json:"scpActionDefinition" yaml:"scpActionDefinition"`
	// ssm_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#ssm_action_definition TfBudgetAction#ssm_action_definition}
	// Experimental.
	SsmActionDefinition *TfBudgetAction_SsmActionDefinitionProperty `field:"optional" json:"ssmActionDefinition" yaml:"ssmActionDefinition"`
}


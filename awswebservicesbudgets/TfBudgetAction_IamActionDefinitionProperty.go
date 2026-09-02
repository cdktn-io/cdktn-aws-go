package awswebservicesbudgets


// Experimental.
type TfBudgetAction_IamActionDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#policy_arn TfBudgetAction#policy_arn}.
	// Experimental.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#groups TfBudgetAction#groups}.
	// Experimental.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#roles TfBudgetAction#roles}.
	// Experimental.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#users TfBudgetAction#users}.
	// Experimental.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}


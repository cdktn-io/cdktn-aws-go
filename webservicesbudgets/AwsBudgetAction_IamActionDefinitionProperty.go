package webservicesbudgets


// Experimental.
type AwsBudgetAction_IamActionDefinitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#policy_arn AwsBudgetAction#policy_arn}.
	// Experimental.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#groups AwsBudgetAction#groups}.
	// Experimental.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#roles AwsBudgetAction#roles}.
	// Experimental.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#users AwsBudgetAction#users}.
	// Experimental.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}


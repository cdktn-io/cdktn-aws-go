package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudgetAction_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#create AwsBudgetsBudgetAction#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#delete AwsBudgetsBudgetAction#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget_action#update AwsBudgetsBudgetAction#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


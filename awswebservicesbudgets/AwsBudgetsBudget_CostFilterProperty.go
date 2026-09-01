package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudget_CostFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#name AwsBudgetsBudget#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#values AwsBudgetsBudget#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


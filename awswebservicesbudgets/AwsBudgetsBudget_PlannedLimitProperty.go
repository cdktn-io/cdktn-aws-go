package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudget_PlannedLimitProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#amount AwsBudgetsBudget#amount}.
	// Experimental.
	Amount *string `field:"required" json:"amount" yaml:"amount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#start_time AwsBudgetsBudget#start_time}.
	// Experimental.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#unit AwsBudgetsBudget#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}


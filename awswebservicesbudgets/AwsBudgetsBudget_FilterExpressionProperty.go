package awswebservicesbudgets


// Experimental.
type AwsBudgetsBudget_FilterExpressionProperty struct {
	// and block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#and AwsBudgetsBudget#and}
	// Experimental.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// cost_categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#cost_categories AwsBudgetsBudget#cost_categories}
	// Experimental.
	CostCategories *AwsBudgetsBudget_FilterExpressionCostCategoriesProperty `field:"optional" json:"costCategories" yaml:"costCategories"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#dimensions AwsBudgetsBudget#dimensions}
	// Experimental.
	Dimensions *AwsBudgetsBudget_FilterExpressionDimensionsProperty `field:"optional" json:"dimensions" yaml:"dimensions"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#not AwsBudgetsBudget#not}
	// Experimental.
	Not *AwsBudgetsBudget_FilterExpressionNotProperty `field:"optional" json:"not" yaml:"not"`
	// or block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#or AwsBudgetsBudget#or}
	// Experimental.
	Or interface{} `field:"optional" json:"or" yaml:"or"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/budgets_budget#tags AwsBudgetsBudget#tags}
	// Experimental.
	Tags *AwsBudgetsBudget_FilterExpressionTagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

